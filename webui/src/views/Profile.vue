<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			ut: null,
			users: null,
			id:null,
			likeStatus: false,
			followStatus: false,
			banStatus: false,
			showComments : false,
			comment : null,
			comments: [],
			chI:false,
			idc: null,
			idp:null,
			salvato: null,
		}
	},
	methods: {
		load() {
			return load
		},
		async getUser() {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/search/" + this.$route.params.username,{
                    	headers: {
							Authorization: localStorage.getItem("token")
						}
                });
				this.users = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getOneUser() {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/search/" + this.$route.params.username,{
                    	headers: {
							Authorization: localStorage.getItem("token")
						}
                });

				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.showComments=false;
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
			this.getOneUser()
			
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
			this.getOneUser()
			if (this.likeStatus==true){
				this.likeStatus=false
			}
			else{
				alert("There is no Like to remove!")
			}
		},

		async followUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put("/users/"+ localStorage.getItem("id").toString()+"/follow/"+id.toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
			if (this.followStatus==false){
				this.followStatus=true
			}
			else{
				alert("User already followed!")
			}
		},
		async unfollowUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/users/"+ localStorage.getItem("id").toString()+"/follow/"+id.toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
			if (this.followStatus==true){
				this.followStatus=false
			}
			else{
				alert("User already unfollowed!")
			}
		},
		async banUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				this.salvato=id
				let response = await this.$axios.put("/users/"+ localStorage.getItem("id").toString()+"/ban/"+id.toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
			if (this.banStatus==false){
				this.banStatus=true
				alert("User banned correctly")
			}
			else{
				alert("User already banned!")
			}
		},
		async unbanUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {

				console.log(this.salvato)
				let response = await this.$axios.delete("/users/"+ localStorage.getItem("id").toString()+"/ban/"+this.salvato.toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
			if (this.banStatus==true){
				this.banStatus=false
				alert("User unbanned correctly")
			}
			else{
				alert("User already unbanned!")
			}
		},
		async goBack() {
			this.$router.push("/stream");
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
			this.comment = ""
			this.getOneUser()
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
			this.getOneUser()
		},
		

	},

	

}
</script>


<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Profile</h1>
			<a href="javascript:" class="btn btn-primary" @click="goBack()">Home</a>
			
		</div>
		<div class="mb-3">
			<a href="javascript:" class="btn btn-info" @click="getOneUser()">Informations</a>
		</div>
		<LoadingSpinner v-if="loading"></LoadingSpinner>



		<div class="card" v-if="!loading" v-for="u in users" :key="u.Id">
		<div class="card-body">
			<a href="javascript:" class="btn btn-warning"  @click="followUser(u.Id)">Follow</a>
			<a href="javascript:" class="btn btn-danger"  @click="unfollowUser(u.Id)">Unfollow</a>
			<p></p>
			<a href="javascript:" class="btn btn-warning" @click="banUser(u.Id)">Ban</a>
			<a href="javascript:" class="btn btn-danger" @click="unbanUser(u.Id)">Unban</a>
		</div>

			<div class="card-body">
				<p class="card-text">
					Username:{{" "  + u.Username }}<br/>
					Numero di foto:{{" "  + u.N_p }}<br/>
					Numero di followers:{{" "  + u.N_followers }}<br/>
					Numero di followings:{{" "  + u.N_followings }}<br/>
				</p>
				<div v-if="!loading" v-for="p in u.Photos" :key="p.Id_photo">
						<img :src="'data:image/png;base64,' + p.Photo" width=300 height=300 /><br/>
						Uploaded: {{p.Data}}<br/>
						Likes:{{p.N_like}}
						Comments:{{p.N_comm}}<br/>
						<a href="javascript:" class="btn btn-primary"  @click="likePhoto(p.Id_photo)">Like</a>
						<a href="javascript:" class="btn btn-danger"  @click="unlikePhoto(p.Id_photo)">Unlike</a>
						<a href="javascript:" class="btn btn-warning" @click="getComments(p.Id_photo)">Comments</a>
						
						<input type="string" class="form-control" id="comment" v-model="comment" placeholder="enter the comment">
						<a href="javascript:" class="btn btn-success" @click="commentPhoto(p.Id_photo)">Send Comment</a>
						

				</div>
			</div>
		</div>
		<div v-if="showComments" class="modal-overlay">
            <div class="modal-content">
                <h2>Comments</h2>

                <div v-for="c in comments" :key="c.Comment_id">
					{{c.Owner_us}}
					comment: {{c.Comment}}
					<a  href="javascript:" class="btn btn-danger" v-if="checkID(c.Owner_id)==true">Delete Comment</a>
					<a  href="javascript:" class="btn btn-danger" v-if="chI==true" style="width: 170px; height: 25px;" @click="deleteComment(c.Comment_id,c.Id_photo)">Delete Comment</a>
					<p></p>
				</div>

				<a href="javascript:" class="btn btn-secondary" @click="getOneUser()">Close</a>
            </div>
        </div>
    </div>
</template>