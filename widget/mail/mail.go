package mail

import (
    cfg "FileServerWeb/config"
    "gopkg.in/gomail.v2"
)


type Mail struct {
    From        string      `json:"from"`
    To          []string    `json:"to"`
    Title       string      `json:"title"`
    Text        string      `json:"text"`
    Text_type   string      `json:"text_type"`
    Attach      string      `json:"attach_path"`
}


func Send_mail(mail Mail) (error) {
    var msg = gomail.NewMessage()
    msg.SetHeader("From", mail.From)
    msg.SetHeader("To", mail.To...)

    msg.SetHeader("Subject", mail.Title)
    msg.SetBody(mail.Text_type, mail.Text)

    if mail.Attach != "" {
        msg.Attach(mail.Attach)
    }

    var dialer = gomail.NewDialer(
        cfg.EMAIL_SMTP_SERVER,
        cfg.EMAIL_SMTP_PORT,
        cfg.EMAIL_USERNAME,
        cfg.EMAIL_PASSWORD,
    )

    var err error
    if err = dialer.DialAndSend(msg); err != nil {
        return err
    }
    return nil
}