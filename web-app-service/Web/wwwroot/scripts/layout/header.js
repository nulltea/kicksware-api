function headerScrollInit() {
	new ScrollMagic.Scene({
		triggerElement: "#scroll-trigger",
		triggerHook: "onLeave"
	}).setClassToggle(".header", "scrolled")
		.addTo( new ScrollMagic.Controller());
}

function changeTheme() {
	let superContainer = $(".super_container");
	if (superContainer.hasClass("theme-light")) {
		superContainer.toggleClass("theme-dark").toggleClass("theme-light");
	} else {
		superContainer.toggleClass("theme-light").toggleClass("theme-dark");
	}
	$(this).toggleClass("dark").toggleClass("light");
}

function themeInit() {
	$(".theme").click(changeTheme);
}

function activePageInit(){
	$(`.main_nav a[href$='${window.location.pathname}']`)
		.parent()
		.toggleClass("active");
}

$(document).ready(function () {
	headerScrollInit();

	themeInit();

	activePageInit();
});