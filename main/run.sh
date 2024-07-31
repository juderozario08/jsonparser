THEN=$(stat -c %Z main.go)
touch main.go
NOW=$(stat -c %Z main.go)

while true
do
    if [[ $THEN != $NOW ]]; then 
        THEN=$NOW
        go run .
    fi
    sleep 0.1
    NOW=$(stat -c %Z main.go)
done
