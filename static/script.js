// script.js

// Function to fetch tasks from the server and display them on the page
function getTasks() {
    fetch('/tasks')
      .then(response => response.json())
      .then(tasks => {
        const taskList = document.getElementById('taskList');
        taskList.innerHTML = '';
  
        tasks.forEach(task => {
          const listItem = document.createElement('li');
          listItem.innerText = `${task.Title} - ${task.Description} - ${task.Status}`;
          taskList.appendChild(listItem);
        });
      })
      .catch(error => console.error('Error fetching tasks:', error));
  }
  
  // Function to add a new task to the server
  function addTask() {
    const taskTitle = document.getElementById('taskTitle').value;
    const taskDescription = document.getElementById('taskDescription').value;
    const taskDueDate = document.getElementById('taskDueDate').value;
    const taskPriority = document.getElementById('taskPriority').value;
    const taskStatus = document.getElementById('taskStatus').value;
  
    fetch('/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        Title: taskTitle,
        Description: taskDescription,
        DueDate: taskDueDate,
        Priority: taskPriority,
        Status: taskStatus,
      }),
    })
    .then(() => {
      alert('New task added successfully!');
      getTasks();
    })
    .catch(error => console.error('Error adding task:', error));
  }
  
  // Function to update an existing task on the server
  function updateTask() {
    const taskIdToUpdate = document.getElementById('taskIdToUpdate').value;
    const updatedTaskTitle = document.getElementById('updatedTaskTitle').value;
    const updatedTaskDescription = document.getElementById('updatedTaskDescription').value;
    const updatedTaskDueDate = document.getElementById('updatedTaskDueDate').value;
    const updatedTaskPriority = document.getElementById('updatedTaskPriority').value;
    const updatedTaskStatus = document.getElementById('updatedTaskStatus').value;
  
    fetch(`/tasks/${taskIdToUpdate}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        Title: updatedTaskTitle,
        Description: updatedTaskDescription,
        DueDate: updatedTaskDueDate,
        Priority: updatedTaskPriority,
        Status: updatedTaskStatus,
      }),
    })
    .then(() => {
      alert('Task updated successfully!');
      getTasks();
    })
    .catch(error => console.error('Error updating task:', error));
  }
  
  // Function to delete an existing task from the server
  function deleteTask() {
    const taskIdToDelete = document.getElementById('taskIdToDelete').value;
  
    fetch(`/tasks/${taskIdToDelete}`, {
      method: 'DELETE',
    })
    .then(() => {
      alert('Task deleted successfully!');
      getTasks();
    })
    .catch(error => console.error('Error deleting task:', error));
  }
  