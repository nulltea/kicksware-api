function changeTheme() {
	let superContainer = $(".super_container");
	if (superContainer.hasClass("theme-light")){
		superContainer.toggleClass("theme-dark").toggleClass("theme-light");
	} else {
		superContainer.toggleClass("theme-light").toggleClass("theme-dark");
	}
	$(this).toggleClass("dark").toggleClass("light");
}


$(document).ready(function () {
	"use strict";

	let header = $(".header");
	let menuActive = false;

	setHeader();

	$(window).on("resize", function () {
		setHeader();
	});

	$(document).on("scroll", function () {
		setHeader();
	});

	$(".theme").click(changeTheme);

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
			window.addEventListener("click", function(event) {
				if (panel.hasClass("active") && !isDescendant(header[0], event.target)) {
					panel.toggleClass("active");
				}
			});
		}
	}

	// INIT MENU //
	function initMenu() {
		var hamButton = $(".hamburger");
		var menu = $(".menu");

		hamButton.on("click", function (event) {
			$(this).toggleClass("open");
			event.stopPropagation();

			if (!menuActive) {
				openMenu();
			}
			else {
				closeMenu();
				menuActive = false;
			}
		});

		window.addEventListener("click", function (event) {
			if (menuActive && !isDescendant(menu[0], event.target)) {
				closeMenu();
				hamButton.toggleClass("open");
			}
		});

		//Handle page menu
		var items = $(".slider-menu-item");
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