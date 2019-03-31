# A Simple Go Client Application For Fetching Open Issues Of Any Github Repository

### Features
The UI displays the following data in a table format when the _owner/organization_ and _repository name_ are entered in the form displayed on the home page:

- Total number of open issues 
- Number of open issues that were opened in the last 24 hours 
- Number of open issues that were opened more than 24 hours ago but less than 7 days ago 
- Number of open issues that were opened more than 7 days ago

### Where The Application Resides

* You can find the [Live Application](https://enigmatic-oasis-10244.herokuapp.com/) by clicking on this link.
* The platform of choice for deploying the application is Heroku.

### Libraries Used

* [gorilla/mux](https://github.com/gorilla/mux) - For routing purposes
* [google/go-github](https://github.com/google/go-github) - Client library for accessing the [Github API V3](https://developer.github.com/v3/)
* [gomodule/redigo/redis](https://github.com/gomodule/redigo) - Client library for communication with Redis
* [streadway/amqp](https://github.com/streadway/amqp) - Client library for communication with RabbitMQ

### Dependency Management

* For dependency management, I am using [golang/dep](https://github.com/golang/dep)

### Approach towards the Solution

* As mentioned, I am using _go-github_ library which does the excellent job of providing a higher level API to access the Github V3 API. 
* Initially, I created a _github client_ object using the structures provided by the _go-github_ package. Since the Github API only allows a certain number of unauthenticated API calls, I created a personal auth token which I use to authenticate the client.
* Next, I passed in various options related to the _features_ discussed above to fetch the required information about a particular _repository_ owned by any _user_.
* After that, I created four routes, one for the _index_ page where a form with two input fields, namely _owner_ and _repository_,  one for creating a worker job where the main _issues_ fetching thing happens, another for checking the _status_ of the completion of the job and the last for actually displaying the data to the user.
* When the relevant information is entered, for example, owner as _smartystreets_ and repository as _goconvey_, the form data is then directed to a _/issues_ API which does the job of calling the actual Github API for fetching the information about the repository by submitting the job to a RabbitMQ queue. This particular job is handled by a background worker which runs as a separate process and it fetches the _job_ to be performed from the RabbitMQ queue.
* To store the data which can be later fetched by the client side, I am using Redis. 
* As some requests take a lot of time process and the request gets stuck, the client sends a request to the server at an interval of 5 seconds to see if the job has been completed or not. If it is done, the client routes the user to the display page where all the results are displayed.

### Areas of Improvement

* I did not particular follow TDD which I should have, but intend to write tests for all the functionalities given more time.
* Also the UI is pretty basic and not very intuitive.
* A *Makefile* would be nice to have as it would ease the setup of project on any computer, thereby supporting the concept of *automation*.
* There is no mechanism present to prevent a large number of requests from any user, which might lead to a server crash as the server is not optimized to handle a large number of requests.
* Implement caching of _issues_ data for a small period of time which would actually save time and resources to fetch the resources from Github.
* Limit the usage of this service by allowing only authenticated clients to make requests and fix a quota for them, so that the API does not get abused.
* Optimize the server to handle a large number of requests and also not accept requests if there is an overloading of requests to it.

**That's all, folks! :D**
