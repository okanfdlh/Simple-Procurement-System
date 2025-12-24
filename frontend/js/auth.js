$("#loginBtn").click(function () {
  $.ajax({
    method: "POST",
    url: API_URL + "/login",
    contentType: "application/json",
    data: JSON.stringify({
      username: $("#username").val(),
      password: $("#password").val()
    }),
    success: function (res) {
      localStorage.setItem("token", res.data.token);
      window.location.href = "dashboard.html";
    },
    error: function () {
      alert("Invalid credentials");
    }
  });
});
