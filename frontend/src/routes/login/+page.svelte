<script>
    import Login from "./Login.svelte";
    import Top from "./../Top.svelte";
    import Footer from "./../Footer.svelte";
</script>
<style>
    #container{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
    }
</style>

<div id="container">
    <Top/>
    <Login/>
    <Footer/>
</div>