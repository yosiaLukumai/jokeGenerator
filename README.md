## JokeGenerator is a simple cli for generating jokes
it illustrate usage of cobra and playing with json, it a silly project i did while tring to consume api's and playing with json.
### Usage
Run the sub-command joke and provide the specific flag for joke type (-t) and count (-c)
>> JokeGenerator joke -c 4 -t "programming"
Fetches a joke basing on any type and number of jokes passed

Running without passing the flag
>> JokeGenerator joke
it will fetch any joke basing on the any type