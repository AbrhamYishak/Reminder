async function send(userData) {
  try {
	  console.log(userData.email)
    const res = await fetch(`${process.env.url}/loginwithgoogle`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ Email: userData.email }),
    });

    const data = await res.json();
    if (res.ok) {
        alert('Login/Registration Successful. Check your email for verification.');
        if (data.redirect === "/setup") {
			localStorage.setItem("ReminderSetupToken", data.token);
            window.location.href = chrome.runtime.getURL("setup.html");
        } else {
			localStorage.setItem("ReminderToken", data.token);
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
	console.log("i am clicked from button")
  chrome.runtime.sendMessage("LOGIN_WITH_GOOGLE", (response) => {
  if (chrome.runtime.lastError) {
    console.error("Runtime error:", chrome.runtime.lastError.message);
    return;
  }
  if (!response) {
    return;
  }
    if (response.error) {
      return;
    }
    const token = response.token;
	console.log(token)
    fetch("https://www.googleapis.com/oauth2/v3/userinfo", {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then(res => res.json())
      .then(data => send(data))
      .catch(err => console.error("Fetch failed:", err));
  });
});
document.addEventListener('DOMContentLoaded',()=> {
  const loginForm = document.getElementById('loginForm');

  if (loginForm) {
    loginForm.addEventListener('submit', async (e)=> {
      e.preventDefault();

      const email = document.getElementById('email').value;
      try {
        const res = await fetch(`${process.env.url}/register`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ Email: email }),
        });

        const data = await res.json();
        if (res.ok) {
          alert('Login Successful. Check your email for verification.');  
		  localStorage.setItem("ReminderSetupToken", data.token);
		  console.log(localStorage.getItem("ReminderSetupToken"))
		  if (data.redirect === "/setup"){
            window.location.href = "setup.html";
			}else{
            alert("Verify your email and open extension")
		  }
        } else {
          alert(data.message || 'Registration failed');
        }
      } catch (err) {
        alert('An error occurred. Please check your network.');
      }
    });
  }
});
