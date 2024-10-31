package main

import (
    "bufio"
    "log"
    "os"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Загрузка цитат из файла
func loadQuotes(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var quotes []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        quotes = append(quotes, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return quotes, nil
}

func main() {
    // Инициализируйте бота с токеном (замените "YOUR_TELEGRAM_BOT_TOKEN" на ваш токен)
    bot, err := tgbotapi.NewBotAPI("7912263108:AAFfTVWtcxHRwXvRRCsMdLoCoqpTsrnNyXY")
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    // Загрузка цитат из файла
    quotes, err := loadQuotes("quotes.txt")
    if err != nil {
        log.Fatalf("Failed to load quotes: %v", err)
    }
    log.Printf("Loaded %d quotes", len(quotes))

    chatID := int64(5838863003) // Замените на ваш chatID

    // Установите более короткий интервал для тестирования (10 секунд)
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    quoteIndex := 0
    for range ticker.C {
        msg := tgbotapi.NewMessage(chatID, quotes[quoteIndex])
        _, err := bot.Send(msg)
        if err != nil {
            log.Printf("Failed to send message: %v", err)
        } else {
            log.Printf("Sent quote: %s", quotes[quoteIndex])
        }

        // Переход к следующей цитате, возврат к началу, если все были отправлены
        quoteIndex = (quoteIndex + 1) % len(quotes)
    }
}

