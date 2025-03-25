
//Contact Form validate  
 $("#contactForm").validate({
  rules: {
      // Define your validation rules here, if any
  },
  messages: {
      // Define custom messages for rules, if any
  },
  submitHandler: function (form) {
      // Create a Bootstrap 5.3 dismissible alert
      const alertHTML = `
          <div class="alert alert-success alert-dismissible fade show mt-3" role="alert">
              Message successfully sent!
              <button type="button" class="btn-close contact-msg" data-bs-dismiss="alert" aria-label="Close">
              <i class="fa-solid fa-xmark"></i> 
              </button>
          </div>
      `;

      // Append the alert to a specific container (e.g., above the form)
      $("#contactForm").before(alertHTML);

      // Optionally clear the form after submission
      form.reset();
  }
});

  if($("#newsletter").length > 0){
    $("#newsletter").validate({
      rules: {
       news_email: {
          required: true
        }
      },

      messages: {
       news_email: {
          required: "Please enter email address",
          email: "Please enter a valid email address"
        }

      }
    });
  }

  // offcanvas menu close and active on click
  document.addEventListener("DOMContentLoaded", function(){
    var myOffcanvas = document.getElementById('offcanvasNavbar2');
    var bsOffcanvas = new bootstrap.Offcanvas(myOffcanvas);
    document.getElementById("offcanvasbutton").addEventListener('click',function (e){
      e.preventDefault();
      e.stopPropagation();
      bsOffcanvas.toggle();
    });
  });

  // offcanvas menu close
  $('.nav-link').on( "click", function(){
    $('.close').click(); 
  }); 

  // back to top
  if($("#backtotop").length > 0){
    var btn = $('#backtotop');
    $(window).scroll(function () {
      if ($(window).scrollTop() > 300) {
        btn.addClass('show');
      } else {
        btn.removeClass('show');
      }
    });
    btn.on('click', function scrollToTop(e) {
      window.scrollTo(0, 0);
    });
  }


  $(document).ready(function(){
    $('.my-slider').slick({
      dots: false,
      infinite: true,
      speed: 500,
      autoplaySpeed: 5000,
      autoplay: true,
      slidesToShow: 4,
      slidesToScroll: 1,
      arrows: true,
      prevArrow: '<button type="button" class="slick-prev">&lt;</button>',
      nextArrow: '<button type="button" class="slick-next">&gt;</button>',
      responsive: [
        {
          breakpoint: 1024,
          settings: {
            slidesToShow: 2,
            slidesToScroll: 1,
          },
        },
        {
          breakpoint: 600,
          settings: {
            slidesToShow: 1,
            slidesToScroll: 1,
          },
        },
      ],
    });
  });
   

  $(document).ready(function(){
    $('.city-slider').slick({
        dots: false,
        infinite: true,
        speed: 500,
        autoplaySpeed: 1500,
        autoplay: true,
        slidesToShow: 4,
        slidesToScroll: 1,
        responsive: [
            {
                breakpoint: 1024,
                settings: {
                    slidesToShow: 2
                }
            },
            {
                breakpoint: 768,
                settings: {
                    slidesToShow: 1
                }
            }
        ],
    });
});

// dark & light mode

var togButton = document.getElementById("btnSwitch");
var currThemeMode = localStorage.getItem("wm_dark");
if(!currThemeMode || currThemeMode == 'false'){
  document.documentElement.setAttribute("data-theme", "wm_light");
  togButton.innerHTML = "<i class='fa fa-moon navtext-color'></i>";
}
else{
  document.documentElement.setAttribute("data-theme", "wm_dark");
  togButton.innerHTML = "<i class='fa fa-sun text-white'></i>";
}
togButton.addEventListener("click", toggle);
function toggle(){
  var currThemeMode = localStorage.getItem("wm_dark");
  if (currThemeMode === null){
     var currThemeMode = "false"
  }
  setTheme(currThemeMode);
}
function setTheme(currThemeMode){
  if(currThemeMode == "true" || currThemeMode == true ){
    document.documentElement.setAttribute("data-theme", "wm_light");
    togButton.innerHTML = "<i class='fa fa-moon navtext-color'></i>";
    localStorage.setItem("wm_dark", currThemeMode ? "false" : "true");
  }
  else{
    document.documentElement.setAttribute("data-theme", "wm_dark");
    togButton.innerHTML = "<i class='fa fa-sun text-white'></i>";
    localStorage.setItem("wm_dark", currThemeMode ? "true" : "false");
  }
}

//Preloader
setTimeout(function(){
   $('.preloader').hide();
   },2000);

  $(document).ready(function() {
    $('.location-select').select2();
});

// Initialize WOW.js
new WOW().init();

if($('#cust_chart').length > 0){
const ctx = document.getElementById('cust_chart');
new Chart(ctx, {
  type: 'line',
  data: {
    labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun"],
    datasets: [{
      label: 'Content 1',
      data:[-20, 10, 0, 5, 10, 20],
      fill: false,
      borderColor: '#FFA805',
      tension: 0.3,
      borderWidth: 1,
      pointRadius: 8, // Increase dot size
      pointHoverRadius: 12, // Increase dot size on hover
      pointBackgroundColor: '#FFA805'
    },
    {
      label: 'Content 2',
      data:[10, 20, 30, 40, 20, 55],
      fill: false,
      borderColor: '#AD2F87',
      tension: 0.3,
      borderWidth: 1,
      pointRadius: 8, // Increase dot size
      pointHoverRadius: 12, // Increase dot size on hover
      pointBackgroundColor: '#AD2F87'
    },
    {
      label: 'Content 3',
      data:[20, 15, 17, 25, 50, 20],
      fill: false,
      borderColor: '#FD7600',
      tension: 0.3,
      borderWidth: 1,
      pointRadius: 8, // Increase dot size
      pointHoverRadius: 12, // Increase dot size on hover,
      pointBackgroundColor: '#FD7600'
    }]
  },
  options:{
    plugins: {
      legend: {
        position: 'bottom', // Move the legend to the bottom 
        align: 'start', // Align the legend to the left
        labels: {
          usePointStyle: true, // Use circular point style
          pointStyle: 'circle', // Explicitly set the point style to circle
          padding: 20, // Adjust spacing between items
          boxWidth: 200, // Adjust circle size
      }
      }
    }
  }
});
}

// Current date

$(document).ready(function () {
  // Get today's date
  var today = new Date();
  var day = today.getDate();
  var monthNames = [
      "Jan", "Feb", "March", "April", "May", "June",
      "July", "Aug", "Sept", "Oct", "Nov", "Dec"
  ];
  var month = monthNames[today.getMonth()];
  // var year = today.getFullYear();

  // Add the correct ordinal suffix to the day
  var suffix = (day % 10 === 1 && day !== 11) ? "st" :
               (day % 10 === 2 && day !== 12) ? "nd" :
               (day % 10 === 3 && day !== 13) ? "rd" : "th";

  var formattedDate = day + suffix + " " + month;

  // Display the formatted date
  $('.current-date').text(formattedDate);
});

// current day

$(document).ready(function () {
  // Get today's day of the week
  var today = new Date();
  var dayNames = ["Sun", "Mon", "Tue", "Wed", "Thurs", "Fri", "Sat"];
  var todayName = dayNames[today.getDay()];

  // Check each element with the class 'today-day'
  $('.today-day').each(function () {
      if ($(this).text() === todayName) {
          $(this).addClass('active'); // Add the 'active' class to the day
          $(this).closest('.weather-card').addClass('active'); // Add 'active' to the parent
      }
  });
});



$(document).ready(function () {
  // Get the current year
  const currentYear = new Date().getFullYear();

  // Set the year in the target element
  $("#currentYear").text(currentYear);
});