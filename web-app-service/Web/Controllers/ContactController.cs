using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Constants;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Models;

namespace Web.Controllers
{
	public class ContactController : Controller
	{
		[ViewData]
		public string HeroCoverPath { get; set; } = $"{Constants.FileStoragePath}/heroes/contact-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "Get in touch with us";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "We are thrilled to meet with you";

		[HttpGet]
		[Route("contacts")]
		[Breadcrumb("Contact", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Contact()
		{
			ViewBag.ContactInfo = GetContactInfo();
			return View();
		}

		[HttpPost]
		public IActionResult Contact(ContactViewModel model)
		{
			return Json(new
			{
				Success = true
			});
		}

		private ContactInfoViewModel[] GetContactInfo()
		{
			return new[]
			{
				new ContactInfoViewModel
				{
					Caption = "Information", Email = "info@kicksware.com", PhoneNumber = "+380 66 50 70 532"
				},
				new ContactInfoViewModel
				{
					Caption = "Partnership", Email = "partner@kicksware.com", PhoneNumber = "+380 66 50 70 532"
				},
				new ContactInfoViewModel
				{
					Caption = "Development", Email = "dev@kicksware.com", PhoneNumber = "+380 66 50 70 532"
				},
			};
		}
	}
}
