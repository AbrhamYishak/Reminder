document.getElementById('setupForm').addEventListener('submit', async (e) => {
  e.preventDefault();

  const timezone = document.getElementById('timezone').value;
  const token = localStorage.getItem("ReminderSetupToken");
  try {
    console.log(token)
    const res = await fetch(`${process.env.url}/setupbefore`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify({ TimeZone: timezone })
    });

    const data = await res.json();
    if (res.ok) {
      localStorage.setItem("ReminderToken", data.token);
      alert('Setup Completed.');
	  alert("Open the Extension")
    } else {
      alert(data.message || 'Setup failed');
    }
  } catch (err) {
    console.error(err);
    alert('An error occurred. Please check your network.');
  }
});
