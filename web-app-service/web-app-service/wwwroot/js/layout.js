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

	var header = $('.header');
	var hambActive = false;
	var menuActive = false;

	setHeader();

	$(window).on('resize', function () {
		setHeader();
	});

	$(document).on('scroll', function () {
		setHeader();
	});

	initSearch();
	initMenu();

    /* 

    2. Set Header

    */

	function setHeader() {
		if ($(window).scrollTop() > 100) {
			header.addClass('scrolled');
		}
		else {
			header.removeClass('scrolled');
		}
	}

    /* 

    4. Init Search

    */

	function initSearch() {
		if ($('.search').length && $('.search_panel').length) {
			var search = $('.search');
			var panel = $('.search_panel');

			search.on('click', function () {
                panel.toggleClass('active');
            });
		}
	}

	/*

	5. Login | Sing Up

	*/

    $("#login-btn").click(function () {
		$("#loginModal").modal("show");
    });

    $("#sign-up-btn").click(function () {
		$("#loginModal").modal("show");
    });

    function isDescendant(parent, child) {
        var node = child.parentNode;
        while (node != null) {
            if (node == parent) {
                return true;
            }
            node = node.parentNode;
        }
        return false;
	}

    window.onclick = function(event) {
        var modal = $("#loginModal");
		if (modal.is(':visible') && !isDescendant(modal[0], event.target)) {
			modal.fadeOut("slow");
            modal.modal("hide");
        }
	};



	/* 

	5. Init Menu

	*/

	function initMenu() {
		if ($('.hamburger').length) {
			var hamb = $('.hamburger');

			hamb.on('click', function (event) {
				event.stopPropagation();

				if (!menuActive) {
					openMenu();

					$(document).one('click', function cls(e) {
						if ($(e.target).hasClass('menu_item')) {
							$(document).one('click', cls);
						}
						else {
							closeMenu();
						}
					});
				}
				else {
					$('.menu').removeClass('active');
					menuActive = false;
				}
			});

			//Handle page menu
			if ($('.page_menu_item').length) {
				var items = $('.page_menu_item');
				items.each(function () {
					var item = $(this);

					item.on('click', function (evt) {
						if (item.hasClass('has-children')) {
							evt.preventDefault();
							evt.stopPropagation();
							var subItem = item.find('> ul');
							if (subItem.hasClass('active')) {
								subItem.toggleClass('active');
								TweenMax.to(subItem, 0.3, { height: 0 });
							}
							else {
								subItem.toggleClass('active');
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
	}

	function openMenu() {
		var fs = $('.menu');
		fs.addClass('active');
		hambActive = true;
		menuActive = true;
	}

	function closeMenu() {
		var fs = $('.menu');
		fs.removeClass('active');
		hambActive = false;
		menuActive = false;
	}
});