/* JS Document */

/******************************

[Table of Contents]

1. Vars and Inits
2. Set Header
4. Init Search
5. Init Menu

******************************/

$(document).ready(function () {
	"use strict";

	/* 

	1. Vars and Inits

	*/

	var header = $(".header");
	var menuActive = false;

	setHeader();

	$(window).on("resize", function () {
		setHeader();
	});

	$(document).on("scroll", function () {
		setHeader();
	});

	initSearch();
	initMenu();

	/* 

	2. Set Header

	*/

	function setHeader() {
		if ($(window).scrollTop() > 100) {
			header.addClass("scrolled");
		}
		else {
			header.removeClass("scrolled");
		}
	}

	/* 

	4. Init Search

	*/

	function initSearch() {
		if ($(".search").length && $(".search_panel").length) {
			var header = $("header");
			var search = $(".search");
			var panel = $(".search_panel");
			var dismiss = $(".search_panel .close-button");
			var searchInput = $(".search_panel input");

			search.on("click", function() {
				panel.toggleClass("active");
				setTimeout(function() {
						searchInput.focus();
				}, 500);
			});
			dismiss.on("click", function() {
				panel.toggleClass("active");
			});
			window.addEventListener("click", function() {
				if (panel.hasClass("active") && !isDescendant(header[0], event.target)) {
					panel.toggleClass("active");
				}
			});
		}
	}

	/*

	5. Login | Sing Up

	*/

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

	$(".account svg").click(singUpShow);

	$("#login-btn").click(loginShow);

	$("#sign-up-btn").click(singUpShow);

	function isDescendant(parent, child) {

		if (child.id === "login" || child.id === "sing-up") return true;

		var node = child.parentNode;
		while (node !== null) {
			if (node == parent) {
				return true;
			}
			node = node.parentNode;
		}
		return false;
	}

	window.onclick = function(event) {
		var modal = $("#loginModal");
		if (modal.is(":visible") && !isDescendant(modal[0], event.target)) {
			modal.fadeOut("slow");
			modal.modal("hide");
		}
	};

	//
	// Init Menu
	//
	function initMenu() {
		var hamButton = $(".hamburger");

		hamButton.on("click", function (event) {
			$(this).toggleClass("open");
			event.stopPropagation();

			if (!menuActive) {
				openMenu();

				$(document).one("click", function cls(e) {
					if ($(e.target).hasClass("menu_item")) {
						$(document).one("click", cls);
					}
					else {
						closeMenu();
					}
				});
			}
			else {
				$(".menu").removeClass("active");
				menuActive = false;
			}
		});

		//Handle page menu
		if ($(".page_menu_item").length) {
			var items = $(".page_menu_item");
			items.each(function () {
				var item = $(this);

				item.on("click", function (evt) {
					if (item.hasClass("has-children")) {
						evt.preventDefault();
						evt.stopPropagation();
						var subItem = item.find("> ul");
						if (subItem.hasClass("active")) {
							subItem.toggleClass("active");
							TweenMax.to(subItem, 0.3, { height: 0 });
						}
						else {
							subItem.toggleClass("active");
							TweenMax.set(subItem, { height: "auto" });
							TweenMax.from(subItem, 0.3, { height: 0 });
						}
					}
					else {
						evt.stopPropagation();
					}
				});
			});
		}
	}

	function openMenu() {
		var menu = $(".menu");
		menu.addClass("active");
		menuActive = true;
	}

	function closeMenu() {
		var menu = $(".menu");
		menu.removeClass("active");
		menuActive = false;
	}
});