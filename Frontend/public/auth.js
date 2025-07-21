async function send(userData) {
  try {
	  console.log(userData.email)
    const res = await fetch("http://localhost:8080/loginwithgoogle", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ Email: userData.email }),
    });

    const data = await res.json();
    if (res.ok) {
		localStorage.setItem("ReminderToken", data.token);
        alert('Login/Registration Successful. Check your email for verification.');
        if (data.redirect === "/setup") {
        window.location.href = chrome.runtime.getURL("setup.html");
        } else {
      const popupUrl = chrome.runtime.getURL('index.html');
      window.location.href = popupUrl;
      if (chrome.tabs && chrome.tabs.query) {
        chrome.tabs.query({ active: true, currentWindow: true }, function(tabs) {
          if (tabs[0] && tabs[0].id) {
            chrome.tabs.remove(tabs[0].id);
          }
        });
     }
	 alert("Open the Extension")
    } 
	}else {
      console.error(" Backend error:", data);
      alert("Registration failed: " + (data.message || "Unknown error"));
    }
  } catch (err) {
    console.error(" Error sending email to backend:", err);
    alert("An error occurred while contacting the server.");
  }
}
document.getElementById("loginWithgoogle").addEventListener("click", () => {
  chrome.runtime.sendMessage("LOGIN_WITH_GOOGLE", (response) => {
  if (!response) {
    return;
  }
    if (response.error) {
      return;
    }

    const token = response.token;
    fetch("https://www.googleapis.com/oauth2/v3/userinfo", {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then(res => res.json())
      .then(data => send(data))
      .catch(err => console.error("Fetch failed:", err));
  });
});
document.addEventListener('DOMContentLoaded', function () {
  const loginForm = document.getElementById('loginForm');

  if (loginForm) {
    loginForm.addEventListener('submit', async function (e) {
      e.preventDefault();

      const email = document.getElementById('email').value;

      try {
        const res = await fetch('http://localhost:8080/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ Email: email }),
        });

        const data = await res.json();
        if (res.ok) {
          alert('Login/Registration Successful. Check your email for verification.');
          window.location.href = "login.html";
        } else {
          alert(data.message || 'Registration failed');
        }
      } catch (err) {
        console.error(err);
        alert('An error occurred. Please check your network.');
      }
    });
  }
});
