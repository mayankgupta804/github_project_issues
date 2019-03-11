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
* [google/go-github](https://github.com/google/go-github) - Go Client library for accessing the [Github API V3](https://developer.github.com/v3/)

### Dependency Management

* For dependency management, I am using [golang/dep](https://github.com/golang/dep)

### Approach towards the Solution

* As mentioned, I am using _go-github_ library which does the excellent job of providing a higher level API to access the Github V3 API. 
* Initially, I created a _github client_ object using the structures provided by the _go-github_ package. I used this client and a simple empty _background context_ as a way to access the _issues_ API.
* Next, I passed in various options related to the _features_ discussed above to fetch the required information about a particular _repository_ owned by any _user_.
* After that, I created two routes, one for the _index_ page where a form with two input fields, namely _owner_ and _repository_, is displayed.
* When the relevant information is entered, for example, owner as _smartystreets_ and repository as _goconvey_, the form data is then directed to a _/issues_ API which does the job of calling the actual Github API for fetching the information about the repository.

### Areas of Improvement

* Since I am using a simple github client context, I am only entitled to a certain number of API calls in a given period of time. If the number of calls to the Github API is more than the intended number, we will not get any response because of _rate limits_ on the API.
* To mitigate, the situtation above is to authenticate the client every time an API call is made.
* However, even with a simple client context, we can make more number of calls in a given period of time, provided we using _caching_ mechanism for storing information that we know wouldn't change in the next couple of days, like we can store the information about _issues opened more than 7 days ago_ in _redis_ cache and put a timeout for clearing the cache every 7 days.
* I did not particular follow TDD which I should have, but intend to write tests for all the functionalities given more time.
* Also the UI is pretty basic and not very intuitive.
* A *Makefile* would be nice to have as it would ease the setup of project on any computer, thereby supporting the concept of *automation*.

### Caveat

Gihub considers *pull requests* also as issues, so when the _issues_ are displayed, you will find the _issues_ containing both the number of  actual _issues_ and _pull requests_.

**That's all, folks! :D**
