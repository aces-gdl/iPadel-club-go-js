"use strict";(self.webpackChunkipadelclub=self.webpackChunkipadelclub||[]).push([[53],{6129:(e,a,t)=>{t.d(a,{A:()=>i});var l=t(5043),n=t(3193),c=t(9190),s=t(8265),r=t(7970),o=t(9722),d=t(579);const i=e=>{let{name:a,value:t,label:i,handleupdate:u}=e;const[h,b]=(0,l.useState)([]);return(0,l.useEffect)((()=>{(async()=>{try{const e=await o.A.get("/v1/catalogs/categories");b(e.data)}catch(e){console.error("Error fetching categories:",e)}})()}),[]),(0,d.jsxs)(n.A,{fullWidth:!0,children:[(0,d.jsx)(c.A,{id:"".concat(a,"-select-label"),children:i}),(0,d.jsx)(s.A,{labelId:"".concat(a,"-select-label"),id:"".concat(a,"-select"),value:t,label:i,onChange:u,name:a,children:h.map((e=>(0,d.jsx)(r.A,{value:e.ID,children:e.description},e.ID)))})]})}},9597:(e,a,t)=>{t.d(a,{A:()=>i});var l=t(5043),n=t(3193),c=t(9190),s=t(8265),r=t(7970),o=t(9722),d=t(579);const i=e=>{let{name:a,value:t,label:i,handleupdate:u}=e;const[h,b]=(0,l.useState)([]);return(0,l.useEffect)((()=>{(async()=>{try{const e=await o.A.get("/v1/catalogs/events");b(e.data)}catch(e){console.error("Error fetching events:",e)}})()}),[]),(0,d.jsxs)(n.A,{fullWidth:!0,children:[(0,d.jsx)(c.A,{id:"".concat(a,"-select-label"),children:i}),(0,d.jsx)(s.A,{labelId:"".concat(a,"-select-label"),id:"".concat(a,"-select"),value:t,label:i,onChange:u,name:a,children:h.map((e=>(0,d.jsx)(r.A,{value:e.ID,children:e.name},e.ID)))})]})}},1952:(e,a,t)=>{t.d(a,{A:()=>A});var l=t(9379),n=t(45),c=t(5043),s=t(6240),r=t(2110),o=t(9958),d=t(4496),i=t(9336),u=t(6494),h=t(579);const b=["border","boxShadow","children","content","contentClass","contentSX","darkTitle","secondary","shadow","sx","title"],x={"& .MuiCardHeader-action":{mr:0}},A=(0,c.forwardRef)(((e,a)=>{let{border:t=!0,boxShadow:c,children:A,content:v=!0,contentClass:g="",contentSX:j={},darkTitle:p,secondary:C,shadow:m,sx:f={},title:I}=e,D=(0,n.A)(e,b);const y=(0,s.A)();return(0,h.jsxs)(r.A,(0,l.A)((0,l.A)({ref:a},D),{},{sx:(0,l.A)({border:t?"1px solid":"none",borderColor:y.palette.primary[200]+25,":hover":{boxShadow:c?m||"0 2px 14px 0 rgb(32 40 45 / 8%)":"inherit"}},f),children:[I&&(0,h.jsx)(o.A,{sx:x,title:p?(0,h.jsx)(d.A,{variant:"h3",children:I}):I,action:C}),I&&(0,h.jsx)(i.A,{}),v&&(0,h.jsx)(u.A,{sx:j,className:g,children:A}),!v&&A]}))}))},6053:(e,a,t)=>{t.r(a),t.d(a,{default:()=>A});var l=t(9379),n=t(5043),c=t(1952),s=t(3193),r=t(9190),o=t(8265),d=t(7970),i=t(9722),u=t(579);const h=e=>{const[a,t]=(0,n.useState)([]);return(0,n.useEffect)((()=>{(async()=>{try{const e=await i.A.get("/v1/catalogs/clubs");t(e.data)}catch(e){console.error("Error fetching clubs:",e)}})()}),[]),(0,u.jsxs)(s.A,{fullWidth:!0,children:[(0,u.jsx)(r.A,{id:"club-select-label",children:"Club"}),(0,u.jsx)(o.A,{labelId:"CategoryL",id:e.name,name:e.name,value:e.value,label:"Club",onChange:e.handleupdate,children:a.map((e=>(0,u.jsx)(d.A,{value:e.ID,children:e.name},e.ID)))})]})};var b=t(9597),x=t(6129);const A=()=>{const[e,a]=(0,n.useState)({ClubID:"1",EventID:"1",CategoryID:"1"}),t=t=>{const{name:n,value:c}=t.target;a((0,l.A)((0,l.A)({},e),{},{[n]:c}))};return(0,u.jsxs)(c.A,{title:"Bienvenido a iPadel Club",minHeigth:"300px",children:[(0,u.jsx)(h,{name:"ClubID",value:e.ClubID,label:"Clubs",handleupdate:t}),(0,u.jsx)(b.A,{name:"EventID",value:e.EventID,label:"Eventos",handleupdate:t}),(0,u.jsx)(x.A,{name:"CategoryID",value:e.CategoryID,label:"Categor\xedas",handleupdate:t})]})}}}]);
//# sourceMappingURL=53.c599a133.chunk.js.map