const controller = new ScrollMagic.Controller();

function headerScrollInit() {
	new ScrollMagic.Scene({
		triggerElement: "#scroll-trigger",
		triggerHook: 0
	}).setClassToggle(".header", "scrolled")
		//.addIndicators()
		.addTo(controller);
}

function changeTheme() {
	let body = $("body");
	let theme;
	if (body.hasClass("theme-light")) {
		body.toggleClass("theme-dark").toggleClass("theme-light");
		theme = "Dark";
	} else {
		body.toggleClass("theme-light").toggleClass("theme-dark");
		theme = "Light";
	}
	$(this).toggleClass("dark").toggleClass("light");

	$.get(`/Profile/SetTheme?theme=${theme}`)
	saveTheme(theme)
	fireThemeEvent();
}

function saveTheme(theme) {
	localStorage.setItem("kicksware.theme", theme);
}

function getTheme() {
	let theme = localStorage.getItem("kicksware.theme");
	if (!theme) {
		$.get("Profile/GetTheme", function (response) {
			theme = response["theme"];
		});
	}
	return (theme ?? "dark").toLowerCase();
}

function setTheme(theme) {
	$("body").removeClass("theme-light theme-dark").addClass(`theme-${theme}`);
	$(".theme").removeClass("light dark").addClass(theme);
	fireThemeEvent();
}

function fireThemeEvent() {
	window.dispatchEvent(new CustomEvent("theme-change", { detail: { theme: $(".theme")[0].classList[1] }}));
}

function themeInit() {
	let theme = $(".theme");
	theme.click(changeTheme);
	setTheme(getTheme());
}

function activePageInit(){
	let baseRoute = window.location.pathname.split("/")[1];
	if (baseRoute){
		$(`.main_nav > ul > li > a[href*='/${window.location.pathname.split("/")[1]}']`)
			.parent()
			.toggleClass("active");
	} else {
		$(`.main_nav> ul > li > a[href$='/']`).parent()
			.toggleClass("active");
	}
}

function handleHeroBrightness(imageSrc, callback) {
	let img = document.createElement("img");
	img.src = imageSrc;
	img.style.display = "none";
	document.body.appendChild(img);
	let colorSum = 0;
	img.onload = function() {
		let canvas = document.createElement("canvas");
		canvas.width = this.width;
		canvas.height = this.height;

		let ctx = canvas.getContext("2d");
		ctx.drawImage(this,0,0);

		let imageData = ctx.getImageData(0,0,canvas.width,canvas.height);
		let data = imageData.data;
		let r,g,b,avg;

		for(let x = 0, len = data.length; x < len; x+=4) {
			r = data[x];
			g = data[x+1];
			b = data[x+2];

			avg = Math.floor((r+g+b)/3);
			colorSum += avg;
		}

		let brightness = Math.floor(colorSum / (this.width*this.height));
		callback(brightness);
	}
}

function mobileResponsiveInit(){
	if (isMobile()){
		$("body").addClass("mobile");
	}
}

$(document).ready(function () {
	headerScrollInit();

	activePageInit();

	themeInit();

	mobileResponsiveInit();
});