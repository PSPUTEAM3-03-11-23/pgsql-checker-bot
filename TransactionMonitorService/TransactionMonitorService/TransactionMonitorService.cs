using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Npgsql;


namespace TransactionMonitorService
{
    class TransactionMonitorService
    {
        private readonly string _connectionString;
        private readonly TimeSpan _pollingInterval;
        private readonly TimeSpan _transactionTimeout;

        public TransactionMonitorService(string connectionString, TimeSpan pollingInterval, TimeSpan transactionTimeout)
        {
            _connectionString = connectionString;
            _pollingInterval = pollingInterval;
            _transactionTimeout = transactionTimeout;
        }

        public async Task StartAsync(CancellationToken cancellationToken)
        {
            while (!cancellationToken.IsCancellationRequested)
            {
                await CheckForLongRunningTransactionsAsync();
                await Task.Delay(_pollingInterval, cancellationToken);
            }
        }

        private async Task CheckForLongRunningTransactionsAsync()
        {
            try
            {
                using (var connection = new NpgsqlConnection(_connectionString))
                {
                    await connection.OpenAsync();

                    string checkTransactionsSql = "SELECT * FROM public.check_long_running_transactions(@max_duration)";

                    using (var command = new NpgsqlCommand(checkTransactionsSql, connection))
                    {
                        command.Parameters.AddWithValue("@max_duration", _transactionTimeout);

                        using (var reader = await command.ExecuteReaderAsync())
                        {
                            while (await reader.ReadAsync())
                            {
                                int pid = reader.GetInt32(reader.GetOrdinal("pid"));
                                TimeSpan duration = (TimeSpan)reader["duration"];
                                string query = reader.GetString(reader.GetOrdinal("query"));

                                // Обработайте найденные таймауты транзакций
                                Console.WriteLine($"Transaction {pid} has been running for {duration}. Query: {query}");
                                // Здесь могут быть дополнительные действия, например, уведомления
                            }
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine($"An error occurred while checking for long running transactions: {ex.Message}");
            }
        }
    }
}
