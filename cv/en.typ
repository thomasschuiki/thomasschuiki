#import "brilliant-CV/template.typ": *
#show: layout

#let data = yaml("./vitae.en.yml")

#cvHeader(hasPhoto: false, align: left)

#cvSection("Professional Experience")
#for e in data.Employment [
  #cvEntry(
      title: [#e.Role],
      society: [#e.Name],
      date: [#e.StartDate - #e.EndDate],
      location: [#e.Location],
      description: list(..e.Details),
      logo: "",
      tags: ()
  )
]

#cvSection("Education")

#for e in data.Education [
  #cvEntry(
      title: [#e.Credential],
      society: [#e.Institution],
      date: [#e.StartDate - #e.EndDate],
      location: [#e.Location],
      description: list([#e.Area]),
      logo: "",
      tags: ()
  )
]

#cvSection("Skills")

#cvSkill(
  type: [Languages],
  info: [German #hBar() English]
)

#cvSkill(
  type: [Software Development],
  info: [Python #hBar() Golang #hBar() Git #hBar() Bash #hBar() Cypress #hBar() MySQL #hBar() MongoDB #hBar() Nginx]
)

#cvSkill(
  type: [Infrastructure],
  info: [Ansible #hBar() Terraform #hBar() Packer #hBar() Docker #hBar() Gitlab CI/CD #hBar() Debian/Ubuntu Linux]
)

#cvSkill(
  type: [IT Administration],
  info: [Atlassian Confluence, JIRA #hBar() Nextcloud #hBar() Gitlab #hBar() AWS #hBar() Microsoft 365]
)

#cvFooter()
