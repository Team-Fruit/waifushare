using System;
using System.Text;
using System.IO;
using System.Net.Http;
using System.Windows.Forms;
using System.Runtime.InteropServices;
using System.Collections.Specialized;
using System.Threading;
using System.Reflection;

namespace WaifuShare {
    class Watcher {
        [DllImport("kernel32.dll")]
        private static extern int GetPrivateProfileString(
            string lpApplicationName,
            string lpKeyName,
            string lpDefault,
            StringBuilder lpReturnedstring,
            int nSize,
            string lpFileName
        );

        private const string MUTEX_NAME = "WaifuShareWatcher";

        private static string GetIniValue(string path, string section, string key) {
            StringBuilder sb = new StringBuilder(256);
            GetPrivateProfileString(
                section,
                key,
                string.Empty,
                sb,
                sb.Capacity,
                path
            );
            return sb.ToString();
        }

        private static async void watcherChanged(
            Object source,
            FileSystemEventArgs e
        ) {
            if (e.ChangeType != WatcherChangeTypes.Created) return;
            try {
                Console.WriteLine(e.Name);
                if ((new FileInfo(e.FullPath)).Length == 0)
                    return;

                string iniPath = Path.Combine(
                    Directory.GetParent(
                        Assembly.GetExecutingAssembly().Location
                    ).ToString(),
                    "configure.ini"
                );

                string server = new StringContent(
                    GetIniValue(iniPath, "configure", "server")
                );

                MultipartFormDataContent form = new MultipartFormDataContent();

                FileStream fs = new FileStream(
                    e.FullPath,
                    FileMode.Open,
                    FileAccess.Read
                );

                form.Add(
                    new StringContent(
                        GetIniValue(iniPath, "configure", "username")
                    ),
                    "username"
                );

                form.Add(
                    new StringContent(
                        GetIniValue(iniPath, "configure", "password")
                    ),
                    "password"
                );

                form.Add(
                    new StringContent(
                        e.Name.Split('-')[1]
                    ),
                    "tweet_id"
                );

                byte[] bs = new byte[fs.Length];
                fs.Read(bs, 0, bs.Length);
                fs.Close();

                form.Add(
                    new ByteArrayContent(bs, 0, bs.Length),
                    "image",
                    e.Name
                );

                HttpClient httpClient = new HttpClient();

                HttpResponseMessage response = await httpClient.PostAsync(
                    server + "/api/v1/image",
                    form
                );

                response.EnsureSuccessStatusCode();
                httpClient.Dispose();

            } catch (FileNotFoundException) {
            }
        }

        private static void Main(string[] Args) {
            Mutex mutex = new Mutex(false, MUTEX_NAME);

            bool hasHandle = false;

            try {
                try {
                    hasHandle = mutex.WaitOne(0, false);
                } catch (AbandonedMutexException) {
                    hasHandle = true;
                }

                if (!hasHandle) {
                    MessageBox.Show("多重起動しようとしています");
                    return;
                }

                FileSystemWatcher fsw = new FileSystemWatcher();

                string iniPath = Path.Combine(
                    Directory.GetParent(
                        Assembly.GetExecutingAssembly().Location
                    ).ToString(),
                    "configure.ini"
                );

                fsw.Path = GetIniValue(iniPath, "configure", "dir");
                fsw.Filter = "Twitter-*.*";
                fsw.NotifyFilter = NotifyFilters.FileName;
                fsw.IncludeSubdirectories = true;
                fsw.Created += new FileSystemEventHandler(watcherChanged);
                fsw.EnableRaisingEvents = true;

                for(;;) Thread.Sleep(1000);
            } finally {
                if(hasHandle)
                    mutex.ReleaseMutex();
            }
        }
    }
}

