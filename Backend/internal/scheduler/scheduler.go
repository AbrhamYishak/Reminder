package scheduler

import (
	"backend/models"
	"backend/db"
	"container/heap"
	"fmt"
	"sync"
	"time"
	"backend/internal"
)
var HLock sync.Mutex
var UpdateChan = make(chan struct{})
var Change = make(chan time.Duration)
 func Scheduler () {
	 db := db.Connection()
    var timer *time.Timer
    for {
        HLock.Lock()
        if H.Len() == 0 {
            HLock.Unlock()
            select {
            case <-UpdateChan:
                continue
            }
        }
        next := (*H)[0]
        waitDuration := time.Until(next.Time)
        if waitDuration < 0 {
            waitDuration = 0
        }
        if timer != nil {
            timer.Stop()
        }
        timer = time.NewTimer(waitDuration)
        HLock.Unlock()

        select {
        case <-timer.C:
            HLock.Lock()
            if H.Len() > 0 && !(*H)[0].Time.After(time.Now()) {
                due := heap.Pop(H).(models.Message)
				var message models.Message
				var inmessage models.InactiveMessage
				var u models.User
				if err:= db.First(&message, due.ID).Error; err!=nil{
					fmt.Println("could not find the message with the given id")
					HLock.Unlock()
				}else{
					if err := db.First(&u,message.UserID).Error; err!=nil{
					fmt.Println("could not find the user with the given id")
					HLock.Unlock()
					}
				 inmessage.Message = message.Message
				 inmessage.Link = message.Link
				 inmessage.Time = message.Time
				 inmessage.UserID = message.UserID
                 HLock.Unlock()
                 fmt.Println("Sending email to:", u.Email)
				 l := fmt.Sprintf("Link : %v", message.Link) 
				 m := fmt.Sprintf("Message : %v", message.Message)
				 if err:= internal.SendMail(m,l, u.Email); err != nil{
					 fmt.Println("could not send the email due to invalid email or internet connection")
					 continue
				 }
                 db.Delete(&due, due.ID)
				 db.AutoMigrate(&inmessage)
				 if err := db.Create(&inmessage).Error; err != nil{
					 fmt.Println("could not transfer row to inactive messages")
				 }
			}
            } else {
                HLock.Unlock()
            }
        case <-UpdateChan:
            continue 
        }
    }
}
