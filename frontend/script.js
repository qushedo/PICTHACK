
function loadJSON(callback) {
    var xhr = new XMLHttpRequest();
    xhr.overrideMimeType("application/json");
    xhr.open('GET', 'data.json', true); 
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status == "200") {
        callback(JSON.parse(xhr.responseText));
      }
    };
    xhr.send(null);
  }
  
  // Функция для отображения данных на странице
  function displayCourses(coursesData) {
    // Получаем контейнер для курсов
    const container = document.getElementById("courses-container");
  
    // Для каждого курса в данных
    coursesData.forEach(course => {
      // Создаем блок курса
      const courseDiv = document.createElement("div");
      courseDiv.classList.add("course");
  
      // Создаем элементы внутри блока курса
      const courseName = document.createElement("h3");
      courseName.textContent = course.name;
      const courseDescription = document.createElement("p");
      courseDescription.textContent = course.description;
      const courseLink = document.createElement("a");
      courseLink.textContent = "Перейти к курсу";
      courseLink.href = course.link;
  
      // Добавляем элементы в блок курса
      courseDiv.appendChild(courseName);
      courseDiv.appendChild(courseDescription);
      courseDiv.appendChild(courseLink);
  
      // Добавляем блок курса в контейнер на странице
      container.appendChild(courseDiv);
    });
  }
  
  // Загружаем данные из JSON файла и отображаем их на странице
  loadJSON(function(response) {
    // response содержит данные из JSON файла
    displayCourses(response);
  });