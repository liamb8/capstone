package main 
 
import ( 
   "log" 
   "os" 
) 
 
func main() { 
   // Tuncates a file to 100 bytes. Anything more than 100 bytes will be deleted.
 
   err := os.Truncate("test2.txt", 100) 
   if err != nil { 
      log.Fatal(err) 
   } 
} 