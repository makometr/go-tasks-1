package task6

// Остановить горутину раньше положенного можно получив сообщение о досрочном завершении через
// - канал (закрытие или сообщение) - task9
// - контекст - task4 через Done, task5 через таймаут
// - errgroup - task3
