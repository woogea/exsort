# exsort

Digitizes the delimiter-separated strings entered from the standard input and sorts them.
It can be used to sort the version numbers.
For example, if the maximum number of digits for each v1.2.3 is 3, 
when 1000 is specified for rank, 1 * 1000 ^ 2 + 2 * 1000 ^ 1 + 3 * 1000 ^ 0 will be calculated and sorted in that order.

# usage
```
$ echo "v1.2.3 
v0.10.2
v1.3.5
v0.2.9" | go run main.go --rank=1000 --column=0 --asc=false --includes="^v[0-9]+[.][0-9]+[.][0-9]$" --reg="[.]"

v1.3.5
v0.10.2
v0.2.9
```

# options

| option | description | default |
| ------ | ----------- | ------- |
|column| sort with this columns | 0 |
|asc | sort order is ascend when this is true | false |
|reg | separator in column | "[.]" |
|delimiter | delimiter for column | " " |
|only-column | print only using column | false |
|rank | max number of digit for each ranks | 1000 |
|includs | only use matched regexp. | "" |
