
var userApi = 'http://127.0.0.1:3000/user'

const App = () => {
    localStorage.setItem("isLogin",false)
    var isLogin = localStorage.getItem("isLogin")
    console.log(isLogin)
    if(isLogin == false) return;
    getTodos(renderTodosTodos)
    handleFormGroup()
}


/// get Movies
getTodos = (call) => {
    var userid = localStorage.getItem("id")

    fetch(todoApi+ "/"+userid)
        .then((res) => res.json())
        .then(call)
}   


// Render Movies
renderTodosTodos = (todos) => {
    var htmls = todos.map((todo) => {
        return `<div class="todo-row" id="${todo.todoID}">
                <div id="${todo.todoID}">${todo.content}</div>
                <div class="icons">
                    <button class="icon done-icon">&#10003;</button>                     
                    <button class="icon edit-icon" onclick="editTodo(event)">&#9998;</button>                     
                    <button   class="icon delete-icon" onclick="deleteTodo(event)">&#10005;</button> 
                </div>
                </div>`
    })
    root.innerHTML = htmls.join('')
}

handleLogin = () => {
    var login = document.getElementById("login")
    login.addEventListener('click', (e) => {
        e.preventDefault()
        var username = document.getElementById("username").value.trim()
        var password = document.getElementById("password").value.trim()
        if(!checkEmpty(username, password)) return;
        CheckUserName(username,password)
    })
}

handleSignUp = () => {
    var signUp = document.getElementById("sign-up")
    signUp.addEventListener('click', (e) => {
        e.preventDefault()
        var loginForm = document.getElementById("login-form")
        var signUp = document.getElementById("register-form")
       loginForm.style.display = 'none'
       signUp.style.display = 'block'
    })
}

handlerPrevLogin = () => {
    var prevLogin = document.getElementById("prev-login")
    prevLogin.addEventListener('click', (e) => {
        e.preventDefault()
        var loginForm = document.getElementById("login-form")
        var signUp = document.getElementById("register-form")
        loginForm.style.display = 'block'
        signUp.style.display = 'none'
    })
}

handleCreate = () => {
    var register = document.getElementById("register")
    register.addEventListener('click', (e) => {
        e.preventDefault()
        var username = document.getElementById("username-singup").value.trim()
        var password = document.getElementById("password-signup").value.trim()
        if (checkEmpty(username, password))
        {
            isExistUserName(username, password)
        }
        return;
     })
}

isExistUserName = (username, password) => {
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      }
    fetch(`${userApi}/${username}`, requestOptions)
    .then(response => response.json())
    .then(res => {
        if(res == username)
        {
            alert("Tên đăng nhập đã tồn tại.!!")
            return;
        }
        else {
            CreateAccount(username, password)
            alert("Bạn đã đăng ký thành công")
        }
    })
}   

CheckPassword = (username,password) => {
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
      
      fetch(`${userApi}/${username}/${password}`, requestOptions)
        .then(response => response.json())
        .then(res => {
            if(res == "")
            {
                alert("Mật khẩu không đúng.!!!")
                return;
            }else
            {
                var loginForm = document.getElementById("login-form")
                loginForm.style.display = 'none'
                loginForm.style.display = 'none'
                localStorage.setItem("isLogin",true)
                localStorage.setItem("id",`${res}`)
                App()
            }
        })
}

CheckUserName = (username,password) => {
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      }
    fetch(`${userApi}/${username}`, requestOptions)
        .then(response => response.json())
        .then(res => {
            if(res != username)
            {
                alert("Tên đăng chưa tồn tại.!!")
                return;
            }
            else {
                CheckPassword(username,password)
            }
    })
}


CreateAccount = (username, password) => {
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
        "username": username,
        "password": password
    });

    var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
    };

    fetch(userApi, requestOptions)
}


checkEmpty = (username, password) => {
    if(username == "" || password == "")
    {
        alert("Username && Password can't be empty ")
        return false
    }
    else
        return true
}
// start 
handleLogin()
handlerPrevLogin()
handleSignUp()
handleCreate()