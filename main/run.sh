THEN=$(stat -f "%m" parser.go)
touch parser.go
NOW=$(stat -f "%m" parser.go)

while true; do
    sleep 0.1
    if [[ $THEN != $NOW ]]; then
        clear
        go run main/main.go
        THEN=$(stat -f "%m" parser.go)
    fi
    NOW=$(stat -f "%m" parser.go)
done
