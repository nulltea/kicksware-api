// AUTH //
function loginShow() {
	$("#auth-content").hide();
	$("#login-content").show();
	$("#login-title").text("Log in");
	$("#login-header-msg").text("Log in to your account to buy, sell, comment, and more.");
	$("#facebook-btn-caption").text("Continue with Facebook");
	$("#google-btn-caption").text("Continue with Google");
	$("#email-btn-caption").text("Log in with Email");
	$("#login-footer-msg").html("Don't have an account? <a id='sing-up'>Sing Up</a>");
	$("#sing-up").css("cursor", "pointer")
		.on("click", singUpShow);
	$("#login-privacy").hide();
	$("#email-btn").on("click", authLoginShow);
	$(".auth-form").attr("action", "/Auth/Login");
	if (!$("#loginModal").is(":visible")) {
		$("#loginModal").modal("show");
	}
}

function singUpShow() {
	$("#auth-content").hide();
	$("#login-content").show();
	$("#login-title").text("Create an Account");
	$("#login-header-msg").text("By creating an account you'll be able to buy, sell, comment, and more");
	$("#facebook-btn-caption").text("Sign up with Facebook");
	$("#google-btn-caption").text("Sign up with Google");
	$("#email-btn-caption").text("Sign up with Email");
	$("#login-footer-msg").html("Already have an account? <a id='login'>Log in</a>");
	$("#login").css("cursor", "pointer").on("click", loginShow);
	$("#login-privacy").show();
	$("#email-btn").on("click", authSingUpShow);
	$(".auth-form").attr("action", "/Auth/SignUp");
	let modal = $("#loginModal");
	if (!modal.is(":visible")) {
		modal.modal("show");
	}
}

function authSingUpShow() {
	$("#auth-title").text("Create an Account");
	$("#auth-btn-caption").text("Sing Up");
	$("#auth-footer-msg").hide();
	$(".auth-checkbox .checkbox_title").text("Sign up for emails from Kicksware");
	$("#auth").css("cursor", "pointer")
		.on("click", loginShow);
	$("#auth-privacy").show();
	$("#login-content").hide();
	$("#auth-content").show();
}

function authLoginShow() {
	$("#auth-title").text("Log in");
	$("#auth-btn-caption").text("Log in");
	$("#auth-footer-msg").show();
	$(".auth-checkbox .checkbox_title").text("Remember me");
	$("#sing-up").css("cursor", "pointer");
	$("#auth-privacy").hide();
	$("#login-content").hide();
	$("#auth-content").show();

}


$(document).ready(function () {
	"use strict";
	$(".account svg").click(singUpShow);

	$("#login-btn").click(loginShow);

	$("#sign-up-btn").click(singUpShow);

	window.addEventListener("click", function (event) {
		let modal = $("#loginModal");
		if (modal.is(":visible") && !isDescendant(modal[0], event.target)) {
			modal.fadeOut("slow");
			modal.modal("hide");
		}
	});

	$("#auth-content button[type=submit]").click(function (event) {
		$.post($(".auth-form").attr("action"), model, function(response) {
			window.location.href = response.redirectUrl;
		});
		event.preventDefault();
	})
});