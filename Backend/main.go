package main

import (
	"container/heap"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type MessageHeap []Message
var hLock sync.Mutex
func (h MessageHeap) Len() int { return len(h) }
func (h MessageHeap) Less(i, j int) bool {
	return h[i].Time.Before(h[j].Time) // min-heap by send time
}
func (h MessageHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MessageHeap) Push(x any) {
	*h = append(*h, x.(Message))
}
var change = make(chan time.Duration)
func (h *MessageHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0:n-1]
	return item
}
type Message struct{
	ID  int64
	Name  string
	Email string
	Message string
	Time time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
var updateChan = make(chan struct{})
var h = &MessageHeap{}
func getMessages(c *gin.Context){
	dsn := "root:12345678@tcp(127.0.0.1:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not connect to database"})
		return
	}
	var m []Message
	if err := db.Find(&m).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}
func createMessage(c *gin.Context){
	dsn := "root:12345678@tcp(127.0.0.1:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var m Message
	if err := c.BindJSON(&m); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong input"})
		return
	}
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not connect to database"})
		return
	}
	result := db.Create(&m) 
	if result.Error != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	}

	hLock.Lock()
	heap.Push(h, m)
	hLock.Unlock()

	select {
	case updateChan <- struct{}{}:
	default:
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully created the message"})
}
func sendMail(message string,address []string){
	m := gomail.NewMessage()
	m.SetHeader("From", "abrhamyishakyifat@gmail.com")
	m.SetHeader("To", address...)
	m.SetHeader("Subject", "Hello from Reminder!")
	m.SetBody("text/html", fmt.Sprintf("<h1>Hello there!</h1><p>%v</p>",message))
	d := gomail.NewDialer("smtp.gmail.com", 587, "abrhamyishakyifat@gmail.com", "empg rnvf hrrs wulx")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}else{
		fmt.Println("Message successfully sent")
	}
 
}
func editMail(c *gin.Context){
	id := c.Param("id")
	dsn := "root:12345678@tcp(127.0.0.1:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not connnect with the database"})
		return} 
	var new_message Message
	if err:=c.BindJSON(&new_message); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"not the write json"})
		return
	}
	var message Message
	if result:=db.First(&message,id).Error; result != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not find the message with the given id"})
		return
	}
	message.Message = new_message.Message
	message.Time = new_message.Time
	message.Email = new_message.Email
	message.Name = new_message.Name
	if result:=db.Save(message).Error; result!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not edit the existing data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully updated the message"})
    
}
func delMail(c *gin.Context){
	id := c.Param("id")
	dsn := "root:12345678@tcp(127.0.0.1:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not connnect with the database"})
		return
	} 
	var m Message
	if err:=db.Delete(m, id).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not delete the message"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully deleted the message"})
}
func main(){
	dsn := "root:12345678@tcp(127.0.0.1:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println(err)
		return}
	var messages []Message
	db.Find(&messages)
	heap.Init(h)
	for _, m := range messages {
		heap.Push(h, m)
	}
	fmt.Println(h)
    go func() {
    var timer *time.Timer
    for {
        hLock.Lock()
        if h.Len() == 0 {
            hLock.Unlock()
            select {
            case <-updateChan:
                continue
            }
        }
        next := (*h)[0]
        waitDuration := time.Until(next.Time)
        if waitDuration < 0 {
            waitDuration = 0
        }
        if timer != nil {
            timer.Stop()
        }
        timer = time.NewTimer(waitDuration)
        hLock.Unlock()

        select {
        case <-timer.C:
            hLock.Lock()
            if h.Len() > 0 && !(*h)[0].Time.After(time.Now()) {
                due := heap.Pop(h).(Message)
				var message Message
				if err:= db.First(&message, due.ID).Error; err!=nil{
					fmt.Println("could not find the message with the given id")
					hLock.Unlock()
				}else{
                 hLock.Unlock()
                 fmt.Println("Sending email to:", message.Email)
                 sendMail(message.Message, strings.Split(message.Email, " "))
                 db.Delete(&due, due.ID)
			}
            } else {
                hLock.Unlock()
            }
        case <-updateChan:
            continue 
        }
    }
}()
	router := gin.Default()
	router.POST("/createMessage", createMessage)
	router.GET("/getMessages", getMessages)
	router.PATCH("/editMessage/:id", editMail)
	router.DELETE("/deleteMessage/:id", delMail)
	router.Run()
}
