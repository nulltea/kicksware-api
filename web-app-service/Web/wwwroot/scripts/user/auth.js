function menuButtonInit() {
	$(".account").off("click").click(function () {
		$.get("/Auth/Auth", function(response) {
			if (response["logged"]) {
				window.location.href = response["redirectUrl"];
				return
			}
			$("#auth-modal").html(response["content"]);
			window.redirectURL = response["redirectUrl"];
			showDialog();
		});
	});
}


function loginOAuthInit() {
	if (!$(".auth-dialog").hasClass("oauth")) {
		return
	}
	$("#login-title").text("Log in");
	$("#login-header-msg").text("Log in to your account to buy, sell, comment, and more.");
	$("#facebook-btn-caption").text("Continue with Facebook");
	$("#google-btn-caption").text("Continue with Google");
	$("#email-btn-caption").text("Log in with Email");
	$("#login-footer-msg").html("Don't have an account? <a id='sing-up'>Sing Up</a>");
	$("#login-privacy").hide();
	$("#email-btn").off("click").click(contentManualShow);
	$(".auth-form").attr("action", "/Auth/Login");
}

function singUpOAuthInit() {
	if (!$(".auth-dialog").hasClass("oauth")) {
		return
	}
	$("#login-title").text("Create an Account");
	$("#login-header-msg").text("By creating an account you'll be able to buy, sell, comment, and more");
	$("#facebook-btn-caption").text("Sign up with Facebook");
	$("#google-btn-caption").text("Sign up with Google");
	$("#email-btn-caption").text("Sign up with Email");
	$("#login-footer-msg").html("Already have an account? <a id='login'>Log in</a>");
	$("#login-privacy").show();
	$("#email-btn").off("click").click(contentManualShow);
	$(".auth-form").attr("action", "/Auth/SignUp");
}

function singUpManualInit() {
	if (!$(".auth-dialog").hasClass("manual")) {
		return
	}
	enableContent("manual")
	$("#auth-title").text("Create an Account");
	$("#auth-btn-caption").text("Sing Up");
	$("#auth-footer-msg").hide();
	$(".auth-checkbox .checkbox_title").text("Sign up for emails from Kicksware");
	$("#auth-privacy").show();
	$("#oauth-content").hide();
	$("#manual-content").show();
	let authForm = $(".auth-form");
	authForm.attr("action", "/Auth/SignUp");
	authForm.find("button[type=submit]").off("click").click(function (event) {
		event.preventDefault();
		onAuthFormSubmit(authForm);
	})
}

function loginManualInit() {
	if (!$(".auth-dialog").hasClass("manual")) {
		return
	}
	$("#auth-title").text("Log in");
	$("#auth-btn-caption").text("Log in");
	$("#auth-footer-msg").show();
	$(".auth-checkbox .checkbox_title").text("Remember me");
	$("#auth-privacy").hide();
	$("#oauth-content").hide();
	$("#manual-content").show();

	let authForm = $(".auth-form");
	authForm.attr("action", "/Auth/Login");
	authForm.find("button[type=submit]").off("click").click(function (event) {
		event.preventDefault();
		onAuthFormSubmit(authForm);
	})
}

function onAuthFormSubmit(authForm) {
	$.post(authForm.attr("action"), authForm.serialize(), function(response) {
		if (!response["success"]) {
			showAuthAlert("error", response["error"]);
			return;
		}
		resetAuthAlert();
		if (response["verifyPending"]) {
			if (response["content"]) {
				$("#auth-modal").html(response["content"]);
			} else {
				$(".auth-dialog").removeClass("login")
					.addClass("locked");
			}
			window.redirectURL = response["redirectUrl"];
			enableContent("verify")
			verifyInit();
			return;
		}
		closeDialog();
	});
}

function contentOAuthShow() {
	enableContent("oauth");
	if ($(".auth-dialog").hasClass("login")) {
		loginOAuthInit();
	} else {
		singUpOAuthInit();
	}
}

function contentManualShow() {
	enableContent("manual");
	if ($(".auth-dialog").hasClass("login")) {
		loginManualInit();
	} else {
		singUpManualInit();
	}
}

function verifyInit(){
	if (!$(".auth-dialog").hasClass("verify")) {
		return
	}
	let content = $("#verify-content");
	let confirmButton = content.find("#confirm-button");
	confirmButton.off("click").click(function () {
		$.get("/Auth/Auth", function(response) {
			if (response["logged"]) {
				closeDialog();
				return
			}
			window.location.href = "/";
		});
	})
	$("#resend-email").off("click").click(function (event) {
		event.preventDefault();
		$.get(this.href, function(response) {
			if (response["success"]) {
				let emailSendMsg = $("#email-send-msg")
				let newMessage = emailSendMsg.text().replace(
					"A verification email was sent to",
					"We've sent another confirmation email to"
				)
				emailSendMsg.text(newMessage);
				$(".login-privacy").detach()
				return
			}
			window.location.href = "/";
		});
	})
}

function enableContent(content) {
	let dialog = $(".auth-dialog")[0];
	if (dialog.classList.contains(content)){
		return;
	}
	dialog.classList.remove("oauth", "manual", "verify")
	dialog.classList.add(content);
	resetAuthAlert();
}

function modalInit() {
	window.addEventListener("click", function (event) {
		let modal = $("#auth-modal");
		if (modal.find(".auth-dialog").hasClass("locked")) {
			return;
		}
		if (modal.is(":visible") && !isDescendant(modal[0], event.target)) {
			closeDialog();
		}
	});

	$(".auth-toggle").click(function () {
		$(".auth-dialog").toggleClass("login");
		if (this.id === "login-toggle") {
			loginOAuthInit();
			loginManualInit();
		} else if (this.id === "sign_up-toggle") {
			singUpOAuthInit();
			singUpManualInit();
		}
		resetAuthAlert();
	})
}

function showDialog() {
	menuButtonInit();
	modalInit();
	verifyInit();
	singUpOAuthInit();
	singUpManualInit();
	verifyInit();
	resetAuthAlert();
	$("#auth-modal").modal("show");
}

function closeDialog() {
	$("#auth-modal").fadeOut("slow").modal("hide");
}

function showAuthAlert(mode, message, lifetime = 5) {
	resetAuthAlert(function () {
		$("#auth-modal .alert-banner")
			.addClass(mode)
			.text(message)
			.addClass("active")
		clearTimeout(window.lifetimeHandler)
		window.lifetimeHandler = window.setTimeout(function () {
			resetAuthAlert();
		}, lifetime * 1000);
	});
}


function resetAuthAlert(callback) {
	let banner = $("#auth-modal .alert-banner");
	if (callback) {
		if (banner.hasClass("active")){
			requestAnimationFrame(function () {
				banner.removeClass("active success error warning").text("");
			})
			window.setTimeout(callback, 500);
		}
		callback();
	} else {
		banner.removeClass("active success error warning").text("");
	}
}

function lock(redirectURL){
	$.get("/Auth/Auth", function(response) {
		if (response["logged"]) {
			window.location.href = response["redirectUrl"];
			return
		}
		let modal = $("#auth-modal");
		modal.html(response["content"]);
		window.redirectURL = redirectURL;
		showDialog();
		let dialog = modal.find(".auth-dialog");
		dialog.addClass("locked");
		dialog.find(".close-button").off("click").click(function () {
			window.location.href = "/";
		})
	})
}

$(document).ready(function () {
	"use strict";

	menuButtonInit();
});

