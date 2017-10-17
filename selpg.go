package main

//includes
import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "io"
)

//types
type selpg_args struct {
  start_page int
  end_page int
  in_filename string
  page_len int
  page_type int
}

var progname string //globals

//don't need prototypes

func main() {
  var sa selpg_args
  progname = os.Args[0] //program name is always the first argument

  flag.Usage = usage

  flag.IntVar(&sa.start_page,"s",-1,"start page defaults to -1")
  flag.IntVar(&sa.end_page,"e",-1,"end page defaults to -1")
  flag.IntVar(&sa.page_len,"l",72,"page length")
  flag.IntVar(&sa.page_type,"f",0,"change page")
  flag.Parse() //give the value to variable

  process_args(&sa)
  process_input(&sa)
}

func process_args(sa *selpg_args) {
   //check the arguments enough or not
   if len(os.Args) <= 2 {
     fmt.Printf("%s: not enough arguments\n",progname)
     usage()
     os.Exit(1)
   }
   //handle 1st arg - start page
   if sa.start_page == -1 && sa.end_page != -1 {
     fmt.Printf("%s: 1st arg should be -sstart_page\n",progname)
     usage()
     os.Exit(2)
   }
   if sa.start_page < 0 {
     fmt.Printf("%s: invalid start page\n", progname)
     usage()
     os.Exit(3)
   }
   //handle 2nd arg - end page
   if sa.start_page != -1 && sa.end_page == -1 {
     fmt.Printf("%s: 2nd arg should be -eend_page\n",progname)
     usage()
     os.Exit(4)
   }
   if sa.end_page < sa.start_page || sa.end_page <= 0 {
     fmt.Printf("%s: invalid end page\n", progname)
     usage()
     os.Exit(5)
   }
   //handle optional args
   if sa.page_len < 1 {
     fmt.Printf("%s: invalid page length", progname)
     usage()
     os.Exit(6)
   }
}

func process_input(sa *selpg_args) {
   if flag.NArg() < 0 {
     scanner := bufio.NewScanner(os.Stdin)
     length := 0
     response := ""
     for scanner.Scan() {
       line := scanner.Text()
       p_len := length / sa.page_len
       line = line + "\n"
       if p_len >= sa.start_page && p_len <= sa.end_page {
         response = response + line
       }
       length++
     }
     fmt.Printf("%s", response)
   } else {
        sa.in_filename = flag.Arg(0)
        f,error := os.Open(sa.in_filename)
        if error != nil {
          fmt.Printf("%s", error)
          os.Exit(7)
        }
        reader := bufio.NewReader(f)
        length := 0
        for true {
          line, _, error := reader.ReadLine()
          if error == io.EOF {
            break
          }
          if error != nil {
            fmt.Printf("%s", error)
            os.Exit(8)
          }
          p_len := length / sa.page_len
          if p_len >= sa.start_page && p_len <= sa.end_page {
            fmt.Printf("%s\n", string(line))
          }
          length++
        }
     }
}

func usage() {
  fmt.Printf("the usage of selpg\n")
  fmt.Printf("USAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ]"
	" [ in_filename ]\n", progname)

}
