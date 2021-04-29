package stringtools

import (
   "fmt"
   "strings"
   "strconv"
)

func Grep (input string, filter string) string {
    var output string
    lines :=strings.Split(string(input),"\n")
    for i := range lines {
       if strings.Contains(lines[i], filter) {
          output = output + lines[i] + "\n"
       }
    }
    return output
}

func RemoveDoubleSpaces(input string) string {

   if len(input)>1 {
      for counter:=0; counter<len(input)-1; counter++ {
          input = strings.Replace(input, "  ", " ", -1)
      }
   }
   return input
}

func SplitSpace (input string) []string {
   return strings.Split(RemoveDoubleSpaces(input), " ")
}

func SplitEnter (input string) []string {
   return strings.Split(input, "\n")
}

func ConvertMacHexToDot(mac string) string {
   // FF:FF:FF:FF:FF:FF > ffff.ffff.ffff
   mac = strings.ToLower(strings.Replace(mac, ":", "", -1))
   mac = fmt.Sprintf("%s.%s.%s", mac[0:4], mac[4:8], mac[8:12])
   return mac
}

func ConvertHexToInt(hex string) int {

   hex = "0x" + hex
   var converted int64
   converted, _ = strconv.ParseInt(hex, 0, 32)

   return int(converted)
}

func Has3dots(iprange string) bool {
  if strings.Count (iprange,".") == 3 { return true } 
  return false
}

func FillWithSpace(input string, num int) string {

   if len(input)<num {
      for counter:=len(input); counter<num; counter++ {
          input = input + " "
      }
   }
   return input
}

func NGrep (input string, filter string) string {
    var output string
    lines :=strings.Split(string(input),"\n")
    for i := range lines {
       if !strings.Contains(lines[i], filter) {
          output = output + lines[i] + "\n"
       }
    }
    return output
}

func SpaceFieldsJoin(str string) string {
    return strings.Join(strings.Fields(str), "")
}

func BeforeDelimiter (mydata string, delimiter string) string {
    return strings.Split(mydata, delimiter)[0]
}

func AfterDelimiter (mydata string, delimiter string) string {
    return strings.Join(strings.Split(mydata, delimiter)[1:], delimiter)
}

func Between (mydata string, after string, before string) string {
    return AfterDelimiter(BeforeDelimiter(mydata, before), after)
}

func StripSpecialCharacters(str string) string {
        b := make([]byte, len(str))
        var bl int
        for i := 0; i < len(str); i++ {
                c := str[i]
                if c >= 32 && c < 127 {
                        b[bl] = c
                        bl++
                }
        }
        return string(b[:bl])
}
