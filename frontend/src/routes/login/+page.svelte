<script>
    import Login from "./Login.svelte";
    import Top from "../../lib/components/Top.svelte";
    import Footer from "../../lib/components/Footer.svelte";
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