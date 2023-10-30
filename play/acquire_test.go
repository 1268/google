package play

import (
   "154.pages.dev/http"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Acquire(t *testing.T) {
   http.No_Location()
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play/"
   var token Refresh_Token
   token.Raw, err = os.ReadFile(home + "token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   var c Client
   if err := c.Token.Refresh(token); err != nil {
      t.Fatal(err)
   }
   time.Sleep(time.Second)
   for _, app := range apps {
      platform := Platforms[fmt.Sprint(app.platform)]
      c.Checkin.Raw, err = os.ReadFile(home + platform + ".bin")
      if err != nil {
         t.Fatal(err)
      }
      if err := c.Checkin.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      fmt.Println(app.id)
      if err := c.Acquire(app.id); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
