using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TransactionMonitorService
{
    public class DatabaseInfo
    {
        public string DB_HOST { get; set; }
        public string DB_USER { get; set; }
        public string DB_PASS { get; set; }
        public string DB_NAME { get; set; }
        public int DB_PORT { get; set; }
    }
}
