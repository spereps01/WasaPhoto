<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			username: null,
			risp: null,
            stream: [],
            users: [],
            likeStatus: false,
            showComments: false,
            chI: false,
		}
	},
	methods: {

		async goBack() {
			this.$router.push("/profile");
		},
        async getStream() {
	
			this.loading = true;
			this.errormsg = null;

            
			try {
				let response = await this.$axios.get("/stream",{
                    	headers: {
							Authorization: localStorage.getItem("token")
						}
                });
                this.showComments= false
				this.stream = response.data;
                
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
        async likePhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put("/photo/"+ id.toString()+"/like/"+localStorage.getItem("id").toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.getStream()
			if (this.likeStatus==false){
				this.likeStatus=true
			}
			else{
				alert("Like already added!")
			}
		},
		async unlikePhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/photo/"+ id.toString()+"/like/"+localStorage.getItem("id").toString());
			
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.getStream()
			if (this.likeStatus==true){
				this.likeStatus=false
			}
			else{
				alert("Like already removed!")
			}
		},
        async commentPhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put("/photo/"+id.toString()+"/comment/"+localStorage.getItem("id"),{
					comment: this.comment
				});
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getStream()
		},
		async getComments(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/photo/"+id.toString()+"/comments");
				this.comments = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.showComments = true;
		},
        async checkID(id){
			if( id.toString() == localStorage.getItem("id")){
				this.chI =true;
			}
			else{
				this.chI=false;
			}

		},
		async deleteComment(idp,idc) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/photo/"+idc.toString()+"/comment/"+idp.toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getStream()
		},
		async Myp(us){
			this.$router.push("/ricerca/profile/"+us);

		}



	},
    mounted() {
		this.getStream()
	}
}
</script>





<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Stream</h1>
			<a href="javascript:" class="btn btn-primary" @click="goBack()">MyProfile</a>
        </div>

        <div class="card" v-if="!loading" v-for="s in stream">
            <button type="button" class="btn btn-sm btn-outline-dark" style="width: 100px; height: 40px;" @click="Myp(s.Username)">
                {{s.Username }}<br/>
            </button>
            <img :src="'data:image/png;base64,' + s.Photo" width=300 height=300 />
			<div class="card-body">
            <a href="javascript:"  class="btn btn-primary" @click="likePhoto(s.Id_photo)">Like</a>
            <a href="javascript:" class="btn btn-danger" @click="unlikePhoto(s.Id_photo)">Unlike</a>
            <a href="javascript:" class="btn btn-warning" @click="getComments(s.Id_photo)">Show Comments</a>
			
			
            <input type="string" class="form-control" id="comment" v-model="comment" placeholder="enter the comment">
            <a href="javascript:" class="btn btn-success" style="width: 160px; height: 35px;" @click="commentPhoto(s.Id_photo)">Send Comment</a>
			</div>
			
        </div>

        <div v-if="showComments" class="modal-overlay">
            <div class="modal-content">
                <h2>Comments</h2>

                <div v-for="c in comments">
					user {{c.Owner_id}}
					comment: {{c.Comment}}
					<a  href="javascript:" class="btn btn-danger" v-if="checkID(c.Owner_id)==true">Delete Comment</a>
					<a  href="javascript:" class="btn btn-danger" v-if="chI==true" style="width: 170px; height: 35px;" @click="deleteComment(c.Comment_id,c.Id_photo)">Delete Comment</a>
					<p></p>
				</div>

				<a href="javascript:" class="btn btn-secondary" @click="getStream()">Close</a>
            </div>
        </div>
    </div>
</template>