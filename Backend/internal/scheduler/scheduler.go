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
				var u models.User
				if err:= db.First(&message, due.ID).Error; err!=nil{
					fmt.Println("could not find the message with the given id")
					HLock.Unlock()
				}else{
					if err := db.First(&u,message.UserID).Error; err!=nil{
					fmt.Println("could not find the user with the given id")
					HLock.Unlock()
					}
                 HLock.Unlock()
                 fmt.Println("Sending email to:", u.Email)
				 l := fmt.Sprintf("Link : %v", message.Link) 
				 m := fmt.Sprintf("Message : %v", message.Message)
				 if err:= internal.SendMail(m,l, u.Email); err != nil{
					 fmt.Println("could not send the email due to invalid email or internet connection")
					 continue
				 }
				internal.ToInactiveMessages(message)
				db.Delete(&message)
				var mess models.Message
				if err := db.Order("Time asc").Limit(1).First(&mess).Error; err == nil{
					heap.Push(H, mess)
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
