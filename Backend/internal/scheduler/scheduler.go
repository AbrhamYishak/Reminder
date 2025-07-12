package scheduler

import (
	"backend/models"
	"backend/db"
	"container/heap"
	"fmt"
	"sync"
	"time"
	"strings"
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
				if err:= db.First(&message, due.ID).Error; err!=nil{
					fmt.Println("could not find the message with the given id")
					HLock.Unlock()
				}else{
                 HLock.Unlock()
                 fmt.Println("Sending email to:", message.Email)
				 if err:= internal.SendMail(message.Message, strings.Split(message.Email, " ")); err != nil{
					 fmt.Println("could not send the email due to invalid email or internet connection")
					 continue
				 }
                 db.Delete(&due, due.ID)
			}
            } else {
                HLock.Unlock()
            }
        case <-UpdateChan:
            continue 
        }
    }
}
