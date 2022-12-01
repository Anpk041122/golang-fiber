var todoApi = 'http://127.0.0.1:3000/todo'
var root = document.getElementById("root")


// Enter code 
Enter = (e) => {
   if (e.key === "Enter") {
        var text =  document.getElementById("input-add").value
        if (checkEmptyLogin(text) == 1) 
        {
            postMovie(text, insertDOM)
            document.getElementById("input-add").value = ""
        }
        document.getElementById("input-add").focus()
    }
}

// Create a new movie
postMovie = (data, callback) => {
    var myHeaders = new Headers();
    myHeaders.append("Authorization", "Basic YWRtaW46MTIzNDU2");
    myHeaders.append("Content-Type", "application/json");
    var id =localStorage.getItem("id")
    var raw = JSON.stringify({
        "done": false,
        "content": data,
        "userid":id
    });
    
    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };
    
    fetch(todoApi, requestOptions)
        .then(res => res.json())
        .then(callback)
}

deleteTodo = (e) => {
    var id = e.target.parentElement.parentElement.id.trim()

    var myHeaders = new Headers();
    myHeaders.append("Authorization", "Basic YWRtaW46MTIzNDU2");

    var requestOptions = {
        method: 'DELETE',
        headers: myHeaders,
        redirect: 'follow'
    };

    fetch(todoApi + "/" + id, requestOptions)
        .then(response => response.json())
        .then(deleteDOM)
}

deleteDOM = (id) => {
    var toRemove = document.getElementById(`${id}`)
    toRemove.remove()
}

getMoive = () => {
    var id = document.getElementById("input-search").value.trim()
    if(id == "")
    {
        alert("Vui lòng nhập id.!")
        document.getElementById('input-search').focus()
    }
    else {
        fetch(`${todoApi}/${id}`)
            .then(res => {
                if (res.status == 200) {
                    return   res.json()
                } else {
                    alert("Request is fail.!!")
                }
            })
            .then(movie => {
                var root = document.getElementById('root-search')
                root.innerHTML += `<li class="movie" id="${movie.id}">
                    ${movie.name} ( ${movie.id} )
                    </li>`
            })
    }
}

/// handle update
editTodo = (e) => {
    var value = e.target.parentElement.parentElement.firstElementChild.textContent
    var id = e.target.parentElement.parentElement.firstElementChild.id
    var divID = e.target.parentElement.parentElement
    HandleUpdate(e, value,id)
    var update = document.querySelector(".todo-button-edit")
    update.onclick = () => {
        var data = document.querySelector(".todo-input-edit").value
        if(data == undefined) return;
        var idLocal = localStorage.getItem("id")
        var myHeaders = new Headers();
        myHeaders.append("Authorization", "Basic YWRtaW46MTIzNDU2");
        myHeaders.append("Content-Type", "application/json");

        var raw = JSON.stringify({
            "done": false,
            "content": data,
            "userid":idLocal    
        });

        var requestOptions = {
            method: 'PUT',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };

        fetch(todoApi + "/" + id, requestOptions)
            .then(response => response.json())
            .then(res => {
                divID.innerHTML = `<div id="${res.id}">${res.content}</div>
                                    <div class="icons">
                                        <button class="icon done-icon">&#10003;</button>                     
                                        <button class="icon edit-icon" onclick="editTodo(event)">&#9998;</button>                     
                                        <button   class="icon delete-icon" onclick="deleteTodo(event)">&#10005;</button> 
                                    </div>`
            })
    }
    return;
}

// HandleUpdate
HandleUpdate = (e,value,id) => {
    var idInput = e.target.parentElement.parentElement
    idInput.style.height = "50.91px"
    idInput.innerHTML = `<div>
                            <input
                                placeholder="Update your item"
                                name="text"
                                class="todo-input-edit"
                            />
                        </div>
                        <div class="icons">
                            <button class="todo-button-edit" >Update</button>
                        </div>`
    var elInput = document.querySelector(".todo-input-edit")
    elInput.focus()
    elInput.onblur = () => {
        var text = document.querySelector(".todo-input-edit").value.trim()
        if(text == "") 
        {
            idInput.innerHTML = `<div id="${id}">${value}</div>
                                <div class="icons">
                                    <button class="icon done-icon">&#10003;</button>                     
                                    <button class="icon edit-icon" onclick="editTodo(event)">&#9998;</button>                     
                                    <button   class="icon delete-icon" onclick="deleteTodo(event)">&#10005;</button> 
                                </div>`
        } 
        return; 
    }
}

insertDOM = (todo) => {
    root.innerHTML += `<div class="todo-row" id="${todo.todoID}">
    <div id="${todo.todoID}">${todo.content}</div>
    <div class="icons">
        <button class="icon done-icon">&#10003;</button>                     
        <button class="icon edit-icon" onclick="editTodo(event)">&#9998;</button>                     
        <button   class="icon delete-icon" onclick="deleteTodo(event)">&#10005;</button> 
    </div>
    </div>`
}

// Handle form group
handleFormGroup = () => {
    var btn = document.getElementById("btn-add")
    // Handle add a new todo
    btn.onclick = () => {

        var text =  document.getElementById("input-add").value
        if (checkEmptyLogin(text) == 1) 
        {
            postMovie(text, insertDOM)
            document.getElementById("input-add").value = ""
        }
        document.getElementById("input-add").focus()
    }
}

checkEmptyLogin = (data) => {
    if(data == "")
    {
        alert("Không được để trường trống")
        return 0
    }   
    return 1
}
