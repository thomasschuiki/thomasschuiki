#import "fontawesome.typ": *
#let lang = sys.inputs.LANG
#let data = yaml("data."+lang+".yml")

/* Personal Information */
#let firstName = data.Personal.FirstName

#let lastName = data.Personal.LastName

#let personalInfo = (
  github: "thomasschuiki",
  gitlab: "thomasschuiki",
  email: data.Contact.email,
  linkedin: "thomas-s-6a743967",
  //custom-1: (icon: "", text: "example", link: "https://example.com"),
  //homepage: "jd.me.org",
  //orcid: "0000-0000-0000-0000",
  //researchgate: "John-Doe",
  //extraInfo: "",
)


/* Language-specific */
// Add your own languages while the keys must match the varLanguage variable
#let headerQuoteInternational = (
  "": [#data.Personal.Tagline],
  "en": [#data.Personal.Tagline],
)

#let cvFooterInternational = (
  "": "Curriculum vitae",
  "en": "Curriculum vitae",
)

#let letterFooterInternational = (
  "": "Cover Letter",
  "en": "Cover Letter",
)

#let nonLatinOverwriteInfo = (
  "customFont": "Heiti SC",
  "firstName": "王道尔",
  "lastName": "",
  // submit an issue if you think other variables should be in this array
)

/* Layout Setting */
#let awesomeColor = "darknight" // Optional: skyblue, red, nephritis, concrete, darknight

#let profilePhoto = "../img/a.jpg" // Leave blank if profil photo is not needed

#let varLanguage = "" // INFO: value must matches folder suffix; i.e "zh" -> "./modules_zh"

#let varEntrySocietyFirst = false // Decide if you want to put your company in bold or your position in bold

#let varDisplayLogo = false // Decide if you want to display organisation logo or not
