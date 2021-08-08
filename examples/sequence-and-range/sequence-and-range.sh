// The above snippet fades the left lasers in from alpha 0 to alpha 1 over 0.5
// beats twice. Once at beat 0 and once more at beat 2.
$ go run sequence-and-range.go
[
    {"_customData":{"_color":[1,0,0,0]},"_time":0,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.2]},"_time":0.1,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.4]},"_time":0.2,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.6]},"_time":0.3,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.8]},"_time":0.4,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,1]},"_time":0.5,"_type":2,"_value":5},
    
    {"_customData":{"_color":[1,0,0,0]},"_time":2,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.2]},"_time":2.1,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.4]},"_time":2.2,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.6]},"_time":2.3,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,0.8]},"_time":2.4,"_type":2,"_value":5},
    {"_customData":{"_color":[1,0,0,1]},"_time":2.5,"_type":2,"_value":5}
]

