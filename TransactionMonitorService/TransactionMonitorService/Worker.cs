using System.IO.Pipes;
using System.Text.Json;

namespace TransactionMonitorService
{
    public class Worker : BackgroundService
    {
        private readonly ILogger<Worker> _logger;

        public Worker(ILogger<Worker> logger)
        {
            _logger = logger;
        }

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            // Запускаем бесконечный цикл для обработки подключений
            while (!stoppingToken.IsCancellationRequested)
            {
                // Создаем новый экземпляр задачи для каждого подключения
                _ = Task.Run(async () =>
                {
                    try
                    {
                        // Здесь задаем максимальное количество экземпляров
                        using (NamedPipeServerStream pipeServer =
                        new NamedPipeServerStream("dbConnectToWriteProcess", PipeDirection.InOut, -1))
                        {
                            Console.WriteLine("Named Pipe server waiting for connection...");
                            await pipeServer.WaitForConnectionAsync(stoppingToken);

                            using (StreamReader sr = new StreamReader(pipeServer))
                            {
                                string receivedData = sr.ReadLine();

                                var data = JsonSerializer.Deserialize<DatabaseInfo>(receivedData);

                                var host = data.DB_HOST;
                                var username = data.DB_USER;
                                var password = data.DB_PASS;
                                var database = data.DB_NAME;
                                var port = data.DB_PORT;


                                var connectionString = $"Host={host};Port={port};Username={username};Password={password};Database={database}";
                                Console.WriteLine(connectionString);

                                var pollingInterval = TimeSpan.FromSeconds(30); // Например, проверяем каждые 30 секунд
                                var transactionTimeout = TimeSpan.FromMinutes(1); // Таймаут транзакции
                                if (ProcedureCreator.CreateProcedur(connectionString))
                                {

                                    var service = new TransactionMonitorService(connectionString, pollingInterval, transactionTimeout);

                                    var cancellationTokenSource = new CancellationTokenSource();
                                    await service.StartAsync(cancellationTokenSource.Token);
                                }
                                else Console.WriteLine("Не удалось связаться с БД");
                            }
                            // Разрываем соединение после обработки
                            pipeServer.Disconnect();
                        }
                    }
                    catch (OperationCanceledException)
                    {
                        Console.WriteLine("Не удалось получить данные: OperationCanceledException");
                    }
                    catch (Exception ex)
                    {
                        _logger.LogError(ex, "Error occurred while working with named pipe.");
                    }
                }, stoppingToken);
                await Task.Delay(100, stoppingToken);
            }
        }
    }
}