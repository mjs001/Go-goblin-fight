package main

import (
  "fmt"
  "math/rand"
  "time"
  "strconv"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  playerHealth := 100
  goblinHealth := 60
  var goblinRoll int
  var playersChoice string

  fmt.Println("You are being attacked by a Goblin! What damage do you take?")
  fmt.Println("Type 'r' to roll the dice")
  fmt.Scan(&playersChoice)

  if playersChoice == "r" {
    goblinRoll = rand.Intn(21)
  }else {
    // for {
    //   if playersChoice == "r"{
    //     break;
    //   } else {
    //       fmt.Println("You did not type 'r'!")
    //       fmt.Println("Type 'r' to roll the dice")
    //       fmt.Scan(&playersChoice)
    //   }
    }
    
  }
  playerHealth -= goblinRoll
  fmt.Println("You take %v damage!", goblinRoll)
  fmt.Println("Your health: %v %T", playerHealth, playerHealth)
  fmt.Println("Goblins Health: %v", goblinHealth)

  //goblinRoll := rand.Intn(21)
  //playerRoll := rand.Intn(21)

  fmt.Println("It is time for you to fight! Type '1' for flame sword and type '2' for ice blast")
  // var iceBlastDamage, flameSwordDamage int8
  fmt.Scan(&playersChoice)
err, i := strconv.ParseInt(playersChoice, 10, 64)
fmt.Println(err, i)
  // for {
  //   if goblinHealth == 0 || playerHealth == 0 {
  //     break;
  //   }else {
  //     if playersChoice == 2 {
  //   iceBlastDamage := rand.Intn(31)
  //   goblinHealth -= iceBlastDamage

  //   fmt.Println("You did %v damage!", iceBlastDamage)
  //   fmt.Println("Goblins Health: %v", goblinHealth)
  //   fmt.Println("Your health: %v %T", playerHealth, playerHealth)
  // }else if playersChoice == 1 {
  //   flameSwordDamage := rand.Intn(25)
  //   goblinHealth -= flameSwordDamage
  //   fmt.Println("You did %v damage!", flameSwordDamage)
  //   fmt.Println("Goblins Health: %v")
  //   fmt.Println("Your health: %v %T", playerHealth, playerHealth)
  // } else {
  //   for {
  //     if playersChoice == 1 || playersChoice == 2 {
  //       break;
  //     } else {
  //       fmt.Println("You did not type '1' or '2'!")
  //       fmt.Println("It is time for you to fight! Type '1' for flame sword and type '2' for ice blast")          
  //       fmt.Scan(&playersChoice)
  //     }
  //   }
  // }

  //   }
  // }
  
}

