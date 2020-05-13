﻿function menuButtonInit() {
	$(".account").click(function () {
		$.get("/Auth/Auth", function(response) {
			if (response["isLoggedIn"]) {
				window.location.href = response["redirectUrl"];
				return
			}
			$("#auth-modal").html(response["content"]);
			window.redirectURL = response["redirectUrl"];
			showDialog();
		})
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
	$("#email-btn").on("click", contentManualShow);
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
	$("#email-btn").on("click", contentManualShow);
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
	authForm.find("button[type=submit]").click(function (event) {
		event.preventDefault();
		$.post(authForm.attr("action"), authForm.serialize(), function(response) {
			$(".auth-dialog").removeClass("login").addClass("locked");
			enableContent("verify")
		});
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
	authForm.find("button[type=submit]").click(function (event) {
		event.preventDefault();
		$.post(authForm.attr("action"), authForm.serialize(), function(response) {
			$(".auth-dialog").removeClass("login").addClass("locked");
			enableContent("verify")
		});
	})
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
	confirmButton.click(function () {
		if (redirectURL) {
			window.location.href = redirectURL;
			return;
		}
		closeDialog();
	})
}

function enableContent(content) {
	let dialog = $(".auth-dialog")[0];
	if (dialog.classList.contains(content)){
		return;
	}
	dialog.classList.remove("oauth", "manual", "verify")
	dialog.classList.add(content);
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
	})
}

function showDialog() {
	menuButtonInit();
	modalInit();
	verifyInit();
	singUpOAuthInit();
	singUpManualInit();
	verifyInit();
	$("#auth-modal").modal("show");
}

function closeDialog() {
	$("#auth-modal").fadeOut("slow").modal("hide");
}


function lock(){
	$.get("/Auth/Auth", function(response) {
		if (response["isLoggedIn"]) {
			window.location.href = response["redirectUrl"];
			return
		}
		let modal = $("#auth-modal");
		modal.html(response["content"]);
		window.redirectURL = response["redirectUrl"];
		showDialog();
		modal.addClass("locked");
	})
}

$(document).ready(function () {
	"use strict";

	menuButtonInit();
});

