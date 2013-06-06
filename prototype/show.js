var show = (function(document,undefined){
    var self ={};
    
    self.slides = {};
    self.index = 0;
    
    self.init = function(){
        var slides = this.slides;
        var title, content, section;
        var container = document.getElementById("slides");
        for(var i = 0; i < slides.length; i++){
                s = slides[i];
                title = document.createElement("h2");
                content = document.createElement("div");
                section = document.createElement("section");
                    
                title.innerHTML = s.title;
                content.innerHTML = s.content;
                    
                section.appendChild(title);
                section.appendChild(content);
                container.appendChild(section);
                
                section.style.webkitAnimationPlayState = "paused";
                section.style.webkitAnimationDuration = "2s";
            
                s.e = section;
                if(i == 0){
                    s.e.style.opacity = 1;
                }
        }
    };
    
    var jump = function(index){
        var last = self.index;
        
        var now = self.slides[last].e;
        var future = self.slides[index].e;
        
        
        
        now.addEventListener("webkitAnimationStart", function(){
            this.style.opacity = 0;
        });
        
        now.addEventListener("webkitTransitionEnd", function(){
            this.style.webkitAnimationName = "";
            this.style.webkitAnimationPlayState = "paused";
        }, false);
        
        future.addEventListener("webkitAnimationStart", function(){
            this.style.opacity = 1;
        });
        
        future.addEventListener("webkitAnimationEnd", function(){
            this.style.webkitAnimationName = "";
            this.style.webkitAnimationPlayState = "paused";
        }, false);
        
        
        future.style.webkitAnimationName = "leftfadeout";
        future.style.webkitAnimationDirection = "reverse";
        future.style.webkitAnimationPlayState ="running";
        
        now.style.webkitAnimationName = "leftfadeout";
        now.style.webkitAnimationPlayState = "running";
                
    };
    
    var input = document.getElementById("panel");
    input.childNodes[1].addEventListener("keydown", function(event){
        if(event.keyCode === 13){
            var v = this.value;
            this.value = "";
            input.style.opacity = 0;
            this.blur();
            jump(parseInt(v));
        }
    });
    
    
    document.body.addEventListener("keydown", function(event){
        var c = event.keyCode;
        if(c === 37 || c === 39 || c === 32)event.preventDefault();
    });
    
    document.body.addEventListener("keyup", function(event){
        var c = event.keyCode;
        
        /*left*/
        if(c === 37){
            jump(self.index-1);
            event.preventDefault();
            return
        }
        
        /*Right or Spacebar*/
        if(c === 39 || c === 32){
            jump(self.index+1);
            event.preventDefault();
            return
        }

	console.log(c);
        
        if(c === 190 && event.shiftKey || c === 186 ){
            var l = input.style.opacity^1;
            
            input.style.opacity = l;
            if(l== 1){
                //input.input.focus();
                input.childNodes[1].focus();
            }
        }
    });
    
    
    
    return self;
}(document));
