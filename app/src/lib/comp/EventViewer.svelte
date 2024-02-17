<script lang="ts">
    import 'animate.css';
    import {fade} from "svelte/transition";
    import type {Alert, Event} from "$lib/pb/user_pb";

    export let alerts: Alert[]
    export let event: Event
</script>

<div id="alert" class="animate__animated animate__headShake animate__infinite" in:fade out:fade>
    <img alt="donate gif" src="https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExZWQ0YTN6enRmZHk2NWhrZmRlazVhaXAyYzljemhkMThsY2lodjg0aSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/LdOyjZ7io5Msw/giphy.gif"/>
    <h1>{donate.user} - {donate.amount} {donate.currency}</h1>
    <h2>{donate.message}</h2>
</div>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Lato:ital,wght@0,900;1,400&display=swap');

    img {
        height: 15em;
    }

    #alert {
        display: flex;
        flex-direction: column;
        align-items: center;
        color: #ff4747;

        word-wrap: anywhere;
        font-family: 'Lato', sans-serif;
    }

    h1 {
        font-weight: bolder;
        font-size: 2em;

        margin: 0.5em auto;
    }

    h2 {
        font-size: 1.5em;
        font-weight: bold;

        margin: 0.5em auto;
    }
</style>