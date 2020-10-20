package main


func replaceSpace(s string) string {
    // s = strings.Replace(s, " ", "%20", -1)
    // return s

    var newStr []byte
    for i:=0;i<len(s);i++ {
        if string(s[i]) == string(" ") {
            newStr = append(newStr, '%')
            newStr = append(newStr, '2')
            newStr = append(newStr, '0')
        } else {
            newStr = append(newStr, s[i])
        }
    }

    return string(newStr)
}
