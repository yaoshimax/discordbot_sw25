package main

import(
   "fmt"
   "time"
   "math/rand"
   "os"
   "os/signal"
   "syscall"
   "strconv"

   "github.com/bwmarrin/discordgo"
)

type roll struct {
   n int
   d int
}

var(

   Token = "Bot {{REPLACE_AUTHENTICATION_TOKEN}}"
   DiceMap = map[string] []roll{
                                "/人間": []roll {roll{n:2,d:0}, roll{n:2, d:0}, roll{n:2,d:0}, roll{n:2, d:0}, roll{n:2,d:0}, roll{n:2, d:0}},
                                "/エルフ": []roll {roll{n:2,d:0}, roll{n:2, d:0}, roll{n:1,d:0}, roll{n:2, d:0}, roll{n:2,d:0}, roll{n:2, d:0}},
                                "/ドワーフ": []roll {roll{n:2,d:6}, roll{n:1, d:0}, roll{n:2,d:0}, roll{n:2, d:0}, roll{n:1,d:0}, roll{n:2, d:6}},
                                "/タビット": []roll {roll{n:1,d:0}, roll{n:1, d:0}, roll{n:1,d:0}, roll{n:2, d:0}, roll{n:2,d:6}, roll{n:2, d:0}},
                                "/ルーンフォーク": []roll {roll{n:2,d:0}, roll{n:1, d:0}, roll{n:2,d:0}, roll{n:2, d:0}, roll{n:2,d:0}, roll{n:1, d:0}},
                                "/ナイトメア": []roll {roll{n:2,d:0}, roll{n:2, d:0}, roll{n:1,d:0}, roll{n:1, d:0}, roll{n:2,d:0}, roll{n:2, d:0}},
                                "/リカント": []roll {roll{n:1,d:0}, roll{n:1, d:3}, roll{n:2,d:0}, roll{n:2, d:0}, roll{n:1,d:6}, roll{n:1, d:0}},
                                "/リルドラケン": []roll {roll{n:1,d:0}, roll{n:2, d:0}, roll{n:2,d:0}, roll{n:2, d:6}, roll{n:1,d:0}, roll{n:2, d:0}},
                                "/グラスランナー": []roll {roll{n:2,d:0}, roll{n:2, d:0}, roll{n:1,d:0}, roll{n:2, d:6}, roll{n:1,d:0}, roll{n:2, d:6}},
                                "/メリア": []roll {roll{n:1,d:0}, roll{n:1, d:0}, roll{n:1,d:0}, roll{n:2, d:6}, roll{n:1,d:0}, roll{n:1, d:0}},
                                "/ティエンス": []roll {roll{n:2,d:0}, roll{n:2, d:0}, roll{n:1,d:0}, roll{n:1, d:3}, roll{n:2,d:0}, roll{n:2, d:3}},
                                "/レプラカーン": []roll {roll{n:2,d:0}, roll{n:1, d:0}, roll{n:2,d:0}, roll{n:2, d:0}, roll{n:2,d:0}, roll{n:2, d:0}},
                               }
)

func main(){

    rand.Seed(time.Now().UnixNano())
    discord, err := discordgo.New(Token)
    discord.Token = Token
    if err != nil {
        fmt.Println("Error logging in")
        fmt.Println(err)
    }

    discord.AddHandler(messageCreate)
    discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
    err = discord.Open()
    if err != nil {
       fmt.Println("error opening connection,", err)
       return
    }

    // Wait here until CTRL-C or other term signal is received.
    fmt.Println("Bot is now running.  Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

   _, ok := DiceMap[m.Content]
	if ok {
      message := createDiceResult(DiceMap[m.Content])
      s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + " "+ m.Content[1:]+"の能力ダイス結果です\n"+message)
	}
}

func createDiceResult(rolllist []roll) (ret string){
   s := []string{"A", "B", "C", "D", "E", "F"}
   for i :=0; i<6; i++{
      ret += s[i]
      ret += ": "
      ret += diceRoll(rolllist[i].n, rolllist[i].d)
      ret += "\n"
   }
   return
}

func diceRoll(n int, d int) (ret string){
   ret = strconv.Itoa(n) + "d("
   result := 0
   for i:=0; i<n; i++ {
      if i!=0{
         ret += ","
      }
      t := rand.Intn(6)+1
      ret+=strconv.Itoa(t)
      result += t
   }
   ret += ")"
   if d !=0 {
      ret += "+" + strconv.Itoa(d)
      result += d
      ret += " => "
   }else{
      ret += " ===> "
   }
   ret += strconv.Itoa(result)
   return
}
