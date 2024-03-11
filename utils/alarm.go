// utils/alarms.go
package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/sarita-growexx/note_with_alarm/models"
)


var alertTriggered = make(map[uint]bool)
var alertTriggeredLock sync.Mutex // Mutex for concurrent map access

func SetAlarmForNotes(notes []*models.Note) {
	for _, note := range notes {
		go func(note *models.Note) {

			alertTriggeredLock.Lock()
            defer alertTriggeredLock.Unlock()

            // Check if the alert has already been triggered for this note
            if alertTriggered[note.ID] {
                return
            }


			deadline := note.Deadline
			deadline = deadline.Truncate(time.Second)

			currentTime := time.Now()
			currentTime = currentTime.Truncate(time.Second)

			remainingTime := deadline.Sub(currentTime)
			
			// fmt.Println("Deadline:", note.Deadline)			
			// fmt.Println("Current Time:", time.Now())
			// fmt.Println("Remaining Time:", remainingTime)
			
			switch {
			case remainingTime <= 0:
				// Already overdue
				fmt.Printf("ALERT: Note '%s' is overdue!\n", note.Title)
				  alertTriggered[note.ID] = true
			case remainingTime <= 30*time.Minute:
				// 30 minutes remaining
				fmt.Printf("ALERT: Note '%s' has 30 minutes remaining.\n", note.Title)
				alertTriggered[note.ID] = true
			case remainingTime <= 1*time.Hour:
				// 1 hour remaining
				fmt.Printf("ALERT: Note '%s' has 1 hour remaining.\n", note.Title)
				alertTriggered[note.ID] = true
			case remainingTime <= 6*time.Hour:
				// 6 hours remaining
				fmt.Printf("ALERT: Note '%s' has 6 hours remaining.\n", note.Title)
				alertTriggered[note.ID] = true
			case remainingTime <= 24*time.Hour:
				// 1 day remaining
				fmt.Printf("ALERT: Note '%s' has 1 day remaining.\n", note.Title)
				alertTriggered[note.ID] = true
			}
		}(note)
	}
}
