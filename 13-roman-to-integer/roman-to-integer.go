func romanToInt(s string) int {
    var sum int
    var index int
    for index<len(s){
        currentSymbol := string(s[index])
        currentValue := symbolValue(currentSymbol)

        if index+1<len(s){
            nextSymbol := string(s[index+1])
            nextValue := symbolValue(nextSymbol)
            if currentValue>=nextValue{
                sum += currentValue
            }else{
                sum -= currentValue
            }
        }else{
            sum += currentValue
        }
        index++
    }
    return sum
}

func symbolValue(s string) int{
    if s=="I"{
        return 1
    }else if s=="V"{
        return 5
    }else if s=="X"{
        return 10
    }else if s=="L"{
        return 50
    }else if s=="C"{
        return 100
    }else if s=="D"{
        return 500
    }else if s=="M"{
        return 1000
    }
    return 0
}