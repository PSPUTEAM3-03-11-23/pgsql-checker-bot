using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Npgsql;

namespace TransactionMonitorService
{
    class ProcedureCreator
    {
        public static bool CreateProcedur(string connectionString)
        {
            var createProcedure = @"
            CREATE OR REPLACE FUNCTION public.check_long_running_transactions(max_duration INTERVAL)
RETURNS TABLE(pid INT, duration INTERVAL, query TEXT) AS
$$
BEGIN
    RETURN QUERY
    SELECT
        pg_stat_activity.pid, -- Явно указываем таблицу для столбца pid
        now() - pg_stat_activity.xact_start AS duration,
        pg_stat_activity.query
    FROM
        pg_stat_activity
    WHERE
        -- Транзакция активна
        pg_stat_activity.xact_start IS NOT NULL AND
        -- Транзакция начата более max_duration назад
        now() - pg_stat_activity.xact_start > max_duration AND
        -- Исключаем саму эту транзакцию
        pg_stat_activity.pid <> pg_backend_pid();
END;
$$ LANGUAGE plpgsql;


        ";

            try
            {
                using (var connection = new NpgsqlConnection(connectionString))
                {
                    connection.Open();

                    // Создание команды с SQL для создания хранимой процедуры
                    using (var command = new NpgsqlCommand(createProcedure, connection))
                    {
                        command.ExecuteNonQuery(); // Выполнение команды
                        Console.WriteLine("Stored procedure has been created successfully.");

                    }
                }
                return true;
            }
            catch (Exception ex)
            {
                Console.WriteLine("An error occurred: " + ex.Message);
                return false;
            }
        }
    }
}
