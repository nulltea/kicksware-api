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
	$("#sing-up").css("cursor", "pointer");
	$("#sing-up").on("click", singUpShow);
	$("#login-privacy").hide();
	$("#email-btn").on("click", authLoginShow);
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
	$("#login").css("cursor", "pointer");
	$("#login").on("click", loginShow);
	$("#email-btn").on("click", authSingUpShow);
	$("#login-privacy").show();
	if (!$("#loginModal").is(":visible")) {
		$("#loginModal").modal("show");
	}
}

function authSingUpShow() {
	$("#auth-title").text("Create an Account");
	$("#auth-btn-caption").text("Sing Up");
	$("#auth-footer-msg").hide();
	$("#notify-footer-msg").show();
	$("#auth").css("cursor", "pointer");
	$("#auth").on("click", loginShow);
	$("#auth-privacy").show();
	$("#login-content").hide();
	$("#auth-content").show();
}

function authLoginShow() {
	$("#auth-title").text("Log in");
	$("#auth-btn-caption").text("Log in");
	$("#auth-footer-msg").show();
	$("#notify-footer-msg").hide();
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
		var modal = $("#loginModal");
		if (modal.is(":visible") && !isDescendant(modal[0], event.target)) {
			modal.fadeOut("slow");
			modal.modal("hide");
		}
	});
});