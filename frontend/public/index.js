document.getElementById("registrationForm").addEventListener("submit", register);

function register(event) {
  event.preventDefault();

  let login = document.getElementById("login").value;
  let password = document.getElementById("password").value;

  console.log("Login:", login);
  console.log("Password:", password);

  fetch(`/register`, { 
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ login, password }),
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`Ошибка регистрации: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    console.log("Регистрация успешна:", data);
  })
  .catch(error => {
    console.error("Ошибка запроса:", error);
  });
}
