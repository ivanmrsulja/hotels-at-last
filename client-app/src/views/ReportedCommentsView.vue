<template>
  <v-container>
    <br />
    <v-flex class="text-center">
      <h1>Reported comments</h1>
    </v-flex>
    <br />
    <br />
    <reported-comments-list></reported-comments-list>
  </v-container>
</template>

<script>
import reportedCommentsList from "@/components/reportedComments/reportedCommentsList.vue";
export default {
  components: { reportedCommentsList },
  name: "reported-comments",
  data() {
    return {
      //
    };
  },
};
</script>
