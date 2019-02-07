$(document).ready(function(){
	
  $('.pag').click( function(){
      $('.pag').removeClass("bg-primary");
      $(this).addClass('bg-primary');
  });

   
  $('#busquedaInformatica, #busquedaNormal, #pdf, #doc, #img, #video, #definicion, #borrarHistorial, #borrarFavoritos, #historial, .pag, #fav').click( function(){
    
    var url = $(this).attr("id");
    var s = $("#search-input").val();
    var pag= $(".bg-primary").attr("id");
    
   	console.log(s);
		console.log(url);
        $.ajax({                        
           type: "POST",                 
           url: "ajax",                     
           data:{
             pag: pag,
			      	value:  s, 
			         type: url},
           success: function(data)             
           {
             //console.log(data)
           // if(data=="false"){
           //   $('#resp').html( '<h5 class="text-danger"> No se encontraron Resultados de Busqueda</h5>' );      
           // } else{
              $('#resp').html(data); 
              
  $('a[href*="http"]:not([href="#"]').click( function(event){

    console.log("a")
    var url = $(this).attr("href");
    var s = $("#search-input").val();
    $.ajax({                        
      type: "POST",                 
      url: "/Historial",                     
      data:{        
        ty:  s, 
        link: url},
      success: function(data)             
      { 
        console.log(data)
       $(location).attr('href',url);
       
         //$('#resp').html(data); 
        },
      });
});
              $( '.icon-star, .icon-star-filled' ).click( function(){
               var ur = $(this).attr("id");
               var na = $(this).attr("name");
               console.log(ur);
               var c;
               if ($(this).hasClass('icon-star')){
                 c='icon-star';
                 $( this ).addClass('select');
               }else if ($(this).hasClass('icon-star-filled')){
                  c='icon-star-filled';
                  $( this ).addClass('active');
 
               
              }
               $.ajax({                        
                 type: "POST",                 
                 url: "ajaxFavoritos",                     
                 data:{
                   link: ur,
                   class: c,
                   theme:  s,
                   title: na,          
                 },
 
                 success: function(data)             
                 {
                   
                   if (data=="icon-star"){
                     $ ( ".select" ).removeClass('icon-star');
                     $( ".select" ).addClass('icon-star-filled');
                     $ ( ".select" ).removeClass('select');
                     alert("url agregada como favorito");
                   }else if (data== "icon-star-filled"){
                     $( ".active" ).removeClass('icon-star-filled');
                     $( ".active" ).addClass('icon-star');
                     $( ".active" ).removeClass('active');
                     alert("url eliminada de favorito");
                   }  
                 },       
               });
              });
              
              $(' #borrarHistorial, #borrarFavoritos').click( function(){
    
                var url = $(this).attr("id");
                    $.ajax({                        
                       type: "POST",                 
                       url: "ajax",                     
                       data:{
                           type: url},
                       success: function(data)                        
                       { console.log("borrado")
                       $('#resp').html(data); }
                      });
                    });               
              $('.contenedor-imagen').on('click', function(){
               $('#modal').modal;
               var ruta_imagen = ($(this).find('img').attr('src'));
               $("#imagen-modal").attr('src', ruta_imagen );
              });
              $('#modal').on('click', function(){
                $('#modal').modal('hide');
              });           
             
          //  }
            
          },
          
      });
       
  });

  
 $('#busquedaInformatica').click();
});

 
   