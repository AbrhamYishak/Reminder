document.getElementById('verifyForm').addEventListener('submit', async function (e) {
  e.preventDefault();
  const token = document.getElementById('token').value;

  try {
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ Token: token }),
    });

    const data = await res.json();
    if (res.ok) {
      alert('Verification Completed');
      console.log(data);
      localStorage.setItem("ReminderToken", data.token);

      if (data.redirect === "/setup") {
        window.location.href = chrome.runtime.getURL("setup.html");
      } else {
        window.location.href = chrome.runtime.getURL("index.html");
      }
    } else {
      alert(data.message || 'Verification failed');
    }
  } catch (err) {
    console.error(err);
    alert('An error occurred. Please check your network.');
  }
});
